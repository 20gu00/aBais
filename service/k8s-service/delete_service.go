package service

import (
	"context"
	"errors"

	"go.uber.org/zap"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// 删除service
func (s *k8sService) DeleteService(client *kubernetes.Clientset, serviceName, namespace string) (err error) {
	err = client.CoreV1().Services(namespace).Delete(context.TODO(), serviceName, metav1.DeleteOptions{})
	if err != nil {
		zap.L().Error("删除Service失败", zap.Error(err))
		return errors.New("删除Service失败, " + err.Error())
	}

	return nil
}
