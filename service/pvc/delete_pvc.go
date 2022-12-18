package pvc

import (
	"context"
	"errors"

	"go.uber.org/zap"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// 删除pvc
func (p *pvc) DeletePvc(client *kubernetes.Clientset, pvcName, namespace string) (err error) {
	err = client.CoreV1().PersistentVolumeClaims(namespace).Delete(context.TODO(), pvcName, metav1.DeleteOptions{})
	if err != nil {
		zap.L().Error("S-DeletePvc 删除Pvc失败", zap.Error(err))
		return errors.New("删除Pvc失败, " + err.Error())
	}

	return nil
}
