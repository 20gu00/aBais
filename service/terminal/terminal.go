package service

import (
	"context"
	"encoding/json"
	"fmt"
	k8sClient "github.com/20gu00/aBais/common/k8s-clientset"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/remotecommand"
	"net/http"
	"time"
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

	// new一个TerminalSession类型的pty实例
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
	// URL:
	// https://192.168.2.100:6443/api/v1/namespaces/default/pods/nginx-wf2-778d88d7c-7rmsk/exec?command=%2Fbin%2Fbash&container=nginx-wf2&stderr=true&stdin=true&stdout=true&tty=true
	// 最基本的客户端 restclient
	req := client.CoreV1().RESTClient().Post().
		Resource("pods").
		Name(podName).
		Namespace(namespace).
		SubResource("exec").
		// 命令的参数
		VersionedParams(&v1.PodExecOptions{
			Container: containerName,
			Command:   []string{"/bin/bash"},
			Stdin:     true,
			Stdout:    true,
			Stderr:    true,
			TTY:       true,
		}, scheme.ParameterCodec)

	// remotecommand提供方法与集群建立长连接,并设置stdin stdout 提供基于 SPDY 协议的 Executor interface，进行和 pod 终端的流的传输
	// 主要实现了http 转 SPDY 添加X-Stream-Protocol-Version相关header 并发送请求
	executor, err := remotecommand.NewSPDYExecutor(conf, "POST", req.URL())
	if err != nil {
		return
	}
	// 建立链接之后从请求的sream中发送、读取数据
	// Stream
	err = executor.StreamWithContext(context.Background(), remotecommand.StreamOptions{
		Stdin:             pty,
		Stdout:            pty,
		Stderr:            pty,
		TerminalSizeQueue: pty,
		Tty:               true,
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

// TerminalMessage定义了终端和容器shell交互内容的格式
// Operation是操作类型
// Data是具体数据内容
// Rows和Cols可以理解为终端的行数和列数，也就是宽、高
type TerminalMessage struct {
	Operation string `json:"operation"`
	Data      string `json:"data"`
	Rows      uint16 `json:"rows"`
	Cols      uint16 `json:"cols"`
}

// 初始化一个websocket.Upgrader类型的对象，用于http协议升级为websocket协议
// 一等公民
var upgrader = func() websocket.Upgrader {
	upgrader := websocket.Upgrader{}
	upgrader.HandshakeTimeout = time.Second * 2
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	return upgrader
}()

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
	_, message, err := t.wsConn.ReadMessage()

	if err != nil {
		zap.L().Info("read message err", zap.Error(err))
		return 0, err
	}
	var msg TerminalMessage
	if err := json.Unmarshal([]byte(message), &msg); err != nil {
		zap.L().Info("read parse message err", zap.Error(err))
		return 0, err
	}

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
