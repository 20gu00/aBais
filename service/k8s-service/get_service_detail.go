package service

import (
	"context"
	"errors"

	"go.uber.org/zap"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// 获取service详情
func (s *k8sService) GetServicetDetail(client *kubernetes.Clientset, serviceName, namespace string) (service *corev1.Service, err error) {
	service, err = client.CoreV1().Services(namespace).Get(context.TODO(), serviceName, metav1.GetOptions{})
	if err != nil {
		zap.L().Error("获取Service详情失败", zap.Error(err))
		return nil, errors.New("获取Service详情失败, " + err.Error())
	}

	return service, nil
}
