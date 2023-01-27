package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	k8sClient "github.com/20gu00/aBais/common/k8s-clientset"

	"github.com/gorilla/websocket"
	"go.uber.org/zap"
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/remotecommand"
)

var Terminal terminal

type terminal struct{}

// 定义websocket的handler方法
func (t *terminal) WsHandler(w http.ResponseWriter, r *http.Request) {
	// 解析form入参，获取参数namespace、podName、containerName
	if err := r.ParseForm(); err != nil {
		return
	}
	namespace := r.Form.Get("namespace")
	podName := r.Form.Get("pod_name")
	containerName := r.Form.Get("container_name")
	cluster := r.Form.Get("cluster")
	zap.L().Info(fmt.Sprintf("exec pod: %s, container: %s, namespace: %s, cluster: %s\n", podName, containerName, namespace, cluster))
	// clientSet
	client, err := k8sClient.K8s.GetK8sClient(cluster)
	if err != nil {
		return
	}
	// 加载k8s配置
	conf, err := rest.InClusterConfig()
	if err != nil {
		conf, err = clientcmd.BuildConfigFromFlags("", k8sClient.K8s.KubeConfMap[cluster])
		if err != nil {
			zap.L().Error("创建k8s配置失败", zap.Error(err))
		}
	}

	// 实例化一个终端new一个TerminalSession类型的pty实例 tty
	// websocket的输入和输出
	pty, err := NewTerminalSession(w, r, nil)
	if err != nil {
		zap.L().Error("获取pty失败", zap.Error(err))
		return
	}

	// pty处理关闭
	defer func() {
		zap.L().Info("关闭session")
		pty.Close()
	}()

	// 初始化pod所在的corev1资源组
	// PodExecOptions struct 包括Container stdout stdout  Command 等结构
	// scheme.ParameterCodec 应该是pod 的GVK （GroupVersion & Kind）之类的
	// URL: https://192.168.2.100:6443/api/v1/namespaces/default/pods/nginx-wf2-778d88d7c-7rmsk/exec?command=%2Fbin%2Fbash&container=nginx-wf2&stderr=true&stdin=true&stdout=true&tty=true
	// 最基本的客户端 restclient 操作pod 请求apiserver
	// 建立个客户端执行exec kubectl exec
	// pod执行命令的方式,这就需要pod namespace container和实际执行cmd
	// 建立请求 需要自己构造restclient
	// exec是post请求
	// GVK GVR 远程执行的命令传递到pod container的终端上运行

	// 组装一个与api-server交互的请求
	req := client.CoreV1().RESTClient().Post().
		Resource("pods").
		Name(podName).
		Namespace(namespace).
		// 子资源exec
		SubResource("exec").
		// 命令的参数 标准输入输出 错误输出,同时也开启终端,最后拿到命令执行返回的数据
		VersionedParams(&v1.PodExecOptions{
			Container: containerName,         // 要执行哪一个容器
			Command:   []string{"/bin/bash"}, // bash命令,也需要配置了开启终端才行
			Stdin:     true,
			Stdout:    true,
			Stderr:    true,
			TTY:       true,
		}, scheme.ParameterCodec)

	// remotecommand提供方法与集群建立长连接,与pod中的容器建立个长链接进行交互,并设置stdin stdout 提供基于 SPDY 协议的 Executor interface，进行和 pod 终端的流的传输
	// 主要实现了http 转 SPDY 添加X-Stream-Protocol-Version相关header 并发送请求
	// kubeconfig 请求方法 请求的url(转换成url)
	// exec是POST请求
	// 发送这个请求给apiserver,类似kubectl exec和apiserver交互,remotecommand包创建一个executor和apiserver建立长连接
	executor, err := remotecommand.NewSPDYExecutor(conf, "POST", req.URL())
	if err != nil {
		return
	}

	// 远程连接 处理和container终端之间的数据读取和写入
	// 基于spdy协议的executor 建立链接之后从请求的stream中发送、读取数据

	// executor和apiserver建立了长连接后,使用流的方式来进行数据的传递,apiserver和kubelet交互,即表明要将命令传递给某个容器,kubelet和底下的容器运行时交互比如docker也就是内置的dockershim交互,让他提供某个容器的流终端,从而可以让excutor流传递数据
	// Stream
	err = executor.StreamWithContext(context.Background(), remotecommand.StreamOptions{
		// reader
		Stdin: pty,
		// writer  可以输出到缓冲区或者从缓冲区读取,也可以用websocket来处理
		Stdout:            pty,
		Stderr:            pty,
		TerminalSizeQueue: pty,  // 终端的伸缩 实现了Next()方法,得到下一次的终端的尺寸
		Tty:               true, // 需要终端
	})

	if err != nil {
		msg := fmt.Sprintf("Exec to pod error! err: %v", err)
		zap.L().Info(msg)
		// 将报错返回出去
		pty.Write([]byte(msg))
		// 标记退出stream流
		pty.Done()
	}
}

const END_OF_TRANSMISSION = "\u0004"

// 终端的信息(操作类型,数据,宽,高)
// TerminalMessage定义了终端和容器shell交互内容的格式(终端格式)
// Operation是操作类型
// Data是具体数据内容
// Rows和Cols可以理解为终端的行数和列数，也就是宽、高
type TerminalMessage struct {
	Operation string `json:"operation"`
	Data      string `json:"data"`
	Rows      uint16 `json:"rows"`
	Cols      uint16 `json:"cols"`
}

// 初始化一个websocket.Upgrader类型的对象，用于http协议升级为websocket协议(http->websocket)
// 一等公民
var upgrader = func() websocket.Upgrader {
	upgrader := websocket.Upgrader{}
	// 握手建立连接的超时时间
	upgrader.HandshakeTimeout = time.Second * 2
	// 检查源头
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	return upgrader
}()

// 终端的会话定义(websocket连接,大小尺寸,时候关闭)
// 定义TerminalSession结构体，实现PtyHandler接口
// wsConn是websocket连接
// sizeChan用来定义终端输入和输出的宽和高
// doneChan用于标记退出终端
type TerminalSession struct {
	wsConn   *websocket.Conn
	sizeChan chan remotecommand.TerminalSize
	doneChan chan struct{}
}

// 该方法用于升级http协议至websocket，并new一个TerminalSession类型的对象返回
// 由于会将http升级为websocket协议，所以需要main重新监听个端口
func NewTerminalSession(w http.ResponseWriter, r *http.Request, responseHeader http.Header) (*TerminalSession, error) {
	conn, err := upgrader.Upgrade(w, r, responseHeader)
	if err != nil {
		return nil, err
	}
	session := &TerminalSession{
		wsConn:   conn,
		sizeChan: make(chan remotecommand.TerminalSize),
		doneChan: make(chan struct{}),
	}
	return session, nil
}

// 关闭doneChan，关闭后触发退出终端
func (t *TerminalSession) Done() {
	close(t.doneChan)
}

// 获取web端是否resize，以及是否退出终端
// 返回此时终端的尺寸
func (t *TerminalSession) Next() *remotecommand.TerminalSize {
	select {
	case size := <-t.sizeChan:
		return &size
	case <-t.doneChan:
		return nil
	}
}

// 用于读取web端的输入，接收web端输入的指令内容(message)
func (t *TerminalSession) Read(p []byte) (int, error) {
	// websocket读取msg
	_, message, err := t.wsConn.ReadMessage()

	if err != nil {
		zap.L().Info("read message err", zap.Error(err))
		return 0, err
	}
	var msg TerminalMessage
	// 将数据反序列化进结构体
	if err := json.Unmarshal([]byte(message), &msg); err != nil {
		zap.L().Info("read parse message err", zap.Error(err))
		return 0, err
	}

	// 终端信息中的操作类型
	switch msg.Operation {
	case "stdin":
		return copy(p, msg.Data), nil
	case "resize":
		t.sizeChan <- remotecommand.TerminalSize{Width: msg.Cols, Height: msg.Rows}
		return 0, nil
	case "ping":
		return 0, nil
	default:
		// zap.String()
		zap.L().Info(fmt.Sprintf("unknown message type %s", msg.Operation))
		// return 0, nil
		return 0, fmt.Errorf("unknown message type '%s'", msg.Operation)
	}
}

// 用于向web端输出，接收web端的指令后，将结果返回出去
func (t *TerminalSession) Write(p []byte) (int, error) {
	// 将结构体信息做序列化(前提是有json标签)
	msg, err := json.Marshal(TerminalMessage{
		Operation: "stdout",
		Data:      string(p),
	})
	if err != nil {
		zap.L().Info("write parse message err", zap.Error(err))
		return 0, err
	}
	if err := t.wsConn.WriteMessage(websocket.TextMessage, msg); err != nil {
		zap.L().Info("write message err", zap.Error(err))
		return 0, err
	}
	return len(p), nil
}

// 用于关闭websocket连接
func (t *TerminalSession) Close() error {
	return t.wsConn.Close()
}
