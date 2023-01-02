package service

import (
	"context"
	"errors"
	"k8s.io/apimachinery/pkg/runtime/schema"

	"go.uber.org/zap"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// 获取secret详情
func (s *secret) GetSecretDetail(client *kubernetes.Clientset, secretName, namespace string) (secret *corev1.Secret, err error) {
	secret, err = client.CoreV1().Secrets(namespace).Get(context.TODO(), secretName, metav1.GetOptions{})
	if err != nil {
		zap.L().Error("S-GetSecretDetail 获取Secret详情失败", zap.Error(err))
		return nil, errors.New("获取Secret详情失败, " + err.Error())
	}

	secret.ManagedFields = []metav1.ManagedFieldsEntry{}
	secret.GetObjectKind().SetGroupVersionKind(schema.GroupVersionKind{
		Kind:    "Secret",
		Version: "v1",
	})
	return secret, nil
}
