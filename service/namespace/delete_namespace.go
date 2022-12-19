package service

import (
	"context"
	"errors"

	"go.uber.org/zap"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// 删除namespace
func (n *namespace) DeleteNamespace(client *kubernetes.Clientset, namespaceName string) (err error) {
	err = client.CoreV1().Namespaces().Delete(context.TODO(), namespaceName, metav1.DeleteOptions{})
	if err != nil {
		zap.L().Error("C-DeleteNamespace 删除Namespace失败", zap.Error(err))
		return errors.New("删除Namespace失败, " + err.Error())
	}

	return nil
}
