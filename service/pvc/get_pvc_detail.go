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

// 获取pvc详情
func (p *pvc) GetPvcDetail(client *kubernetes.Clientset, pvcName, namespace string) (pvc *corev1.PersistentVolumeClaim, err error) {
	pvc, err = client.CoreV1().PersistentVolumeClaims(namespace).Get(context.TODO(), pvcName, metav1.GetOptions{})
	if err != nil {
		zap.L().Error("S-GetPvcDetail 获取Pvc详情失败", zap.Error(err))
		return nil, errors.New("获取Pvc详情失败, " + err.Error())
	}

	pvc.ManagedFields = []metav1.ManagedFieldsEntry{}
	pvc.GetObjectKind().SetGroupVersionKind(schema.GroupVersionKind{
		Kind:    "PersistentVolumeClaim",
		Version: "v1",
	})
	return pvc, nil
}
