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

// 获取pv详情
func (p *pv) GetPvDetail(client *kubernetes.Clientset, pvName string) (pv *corev1.PersistentVolume, err error) {
	pv, err = client.CoreV1().PersistentVolumes().Get(context.TODO(), pvName, metav1.GetOptions{})
	if err != nil {
		zap.L().Error("S-GetPvDetail 获取Pv详情失败", zap.Error(err))
		return nil, errors.New("获取Pv详情失败, " + err.Error())
	}

	pv.ManagedFields = []metav1.ManagedFieldsEntry{}
	pv.GetObjectKind().SetGroupVersionKind(schema.GroupVersionKind{
		Kind:    "PersistentVolume",
		Version: "v1",
	})
	return pv, nil
}
