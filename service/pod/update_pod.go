package pod

import (
	"context"
	"encoding/json"
	"errors"

	"go.uber.org/zap"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// 更新pod
func (p *pod) UpdatePod(client *kubernetes.Clientset, podName, namespace, content string) (err error) {
	var pod = &corev1.Pod{}
	// content请求中传入的pod对象的json数据 反序列化为pod对象(格式)
	err = json.Unmarshal([]byte(content), pod)
	if err != nil {
		zap.L().Error("S-UpdatePod 反序列化失败, ", zap.Error(err))
		return errors.New("反序列化失败, " + err.Error())
	}

	_, err = client.CoreV1().Pods(namespace).Update(context.TODO(), pod, metav1.UpdateOptions{})
	if err != nil {
		zap.L().Error("S-UpdatePod 更新Pod失败, ", zap.Error(err))
		return errors.New("更新Pod失败, " + err.Error())
	}

	return nil
}
