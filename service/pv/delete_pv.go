package service

import (
	"context"
	"errors"
	"go.uber.org/zap"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// 删除pv
func (p *pv) DeletePv(client *kubernetes.Clientset, pvName string) (err error) {
	err = client.CoreV1().PersistentVolumes().Delete(context.TODO(), pvName, metav1.DeleteOptions{})
	if err != nil {
		zap.L().Error("S-DeletePv 删除Pv失败", zap.Error(err))
		return errors.New("删除Pv失败, " + err.Error())
	}

	return nil
}
