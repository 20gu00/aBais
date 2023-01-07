package service

import (
	"context"
	"errors"

	"go.uber.org/zap"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/kubernetes"
)

// yaml
func (j *sa) GetSaDetail(client *kubernetes.Clientset, saName, namespace string) (sa *corev1.ServiceAccount, err error) {
	sa, err = client.CoreV1().ServiceAccounts(namespace).Get(context.TODO(), saName, metav1.GetOptions{})
	if err != nil {
		zap.L().Error("S-GetSaDetail 获取sa详情失败", zap.Error(err))
		return nil, errors.New("获取sa详情失败" + err.Error())
	}

	sa.ManagedFields = []metav1.ManagedFieldsEntry{}
	sa.GetObjectKind().SetGroupVersionKind(schema.GroupVersionKind{
		Kind:    "ServiceAccount",
		Version: "core/v1",
	})
	return sa, nil
}
