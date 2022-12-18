package service

import (
	"context"
	"errors"

	"go.uber.org/zap"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// 删除ingress
func (i *ingress) DeleteIngress(client *kubernetes.Clientset, ingressName, namespace string) (err error) {
	err = client.NetworkingV1().Ingresses(namespace).Delete(context.TODO(), ingressName, metav1.DeleteOptions{})
	if err != nil {
		zap.L().Error("删除Ingress失败", zap.Error(err))
		return errors.New("删除Ingress失败, " + err.Error())
	}

	return nil
}
