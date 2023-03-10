package service

import (
	"bytes"
	"context"
	"errors"
	"io"

	"github.com/20gu00/aBais/common/config"

	"go.uber.org/zap"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
)

// 获取pod内容器日志
func (p *pod) GetPodLog(client *kubernetes.Clientset, containerName, podName, namespace string) (log string, err error) {
	//设置日志的配置，容器名、tail的行数
	lineLimit := int64(config.Config.PodLogTailLine)
	// 设置pog的日志option
	option := &corev1.PodLogOptions{
		// Follow -f 持续输出,末尾 要实时,那么下面传递数据给前端就得使用websocket来传递了
		Container: containerName,
		// --tail 末尾多少行
		TailLines: &lineLimit,
	}
	// 获取request实例,日志查询配置  GetLogs
	// -c 指定容器 --all
	req := client.CoreV1().Pods(namespace).GetLogs(podName, option)
	// 发起request请求，request的流,返回一个io.ReadCloser类型 带closer的reader（等同于response.body）
	podLogs, err := req.Stream(context.TODO())
	if err != nil {
		zap.L().Error("获取Pod Log失败", zap.Error(err))
		return "", errors.New("获取Pod Log失败, " + err.Error())
	}
	// 关闭reader
	defer podLogs.Close()

	// 将response body拷贝到缓冲区，在转成string返回
	// 取出podLogs的数据取出来
	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, podLogs)
	if err != nil {
		zap.L().Error("复制Pod Log失败", zap.Error(err))
		return "", errors.New("复制Pod Log失败, " + err.Error())
	}

	// buf->string
	return buf.String(), nil

	//buf:=bufio.NewReader(podLogs)
	//for {
	//	// 按行读取
	//	bytes,err:=buf.ReadBytes("\n")
	//	if err!=nil{
	//		// 读取完毕
	//		if err!=io.EOF{
	//			panic(err.Error())
	//		}
	//		return
	//	}
	//}
	// websocket传递
}
