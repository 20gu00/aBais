package pod

import (
	"context"
	"errors"

	"go.uber.org/zap"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

//删除pod
func (p *pod) DeletePod(client *kubernetes.Clientset, podName, namespace string) (err error) {
	err = client.CoreV1().Pods(namespace).Delete(context.TODO(), podName, metav1.DeleteOptions{})
	if err != nil {
		zap.L().Error("S-DeletePod 删除pod失败, ", zap.Error(err))
		return errors.New("删除pod失败: " + err.Error())
	}

	return nil
}
