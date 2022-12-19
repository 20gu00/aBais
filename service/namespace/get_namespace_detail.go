package service

import (
	"context"
	"errors"
	"go.uber.org/zap"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// 获取namespace详情
func (n *namespace) GetNamespaceDetail(client *kubernetes.Clientset, namespaceName string) (namespace *corev1.Namespace, err error) {
	namespace, err = client.CoreV1().Namespaces().Get(context.TODO(), namespaceName, metav1.GetOptions{})
	if err != nil {
		zap.L().Error("S-GetNamespaceDetail 获取Namespace详情失败", zap.Error(err))
		return nil, errors.New("获取Namespace详情失败, " + err.Error())
	}

	return namespace, nil
}
