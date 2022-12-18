package service

import (
	"context"
	"errors"

	"go.uber.org/zap"
	nwv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

//获取ingress详情
func (i *ingress) GetIngresstDetail(client *kubernetes.Clientset, ingressName, namespace string) (ingress *nwv1.Ingress, err error) {
	ingress, err = client.NetworkingV1().Ingresses(namespace).Get(context.TODO(), ingressName, metav1.GetOptions{})
	if err != nil {
		zap.L().Error("获取Ingress详情失败", zap.Error(err))
		return nil, errors.New("获取Ingress详情失败, " + err.Error())
	}

	return ingress, nil
}
