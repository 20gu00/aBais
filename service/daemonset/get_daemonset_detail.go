package service

import (
	"context"
	"errors"

	"go.uber.org/zap"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// daemonset详情
func (d *daemonSet) GetDaemonSetDetail(client *kubernetes.Clientset, daemonSetName, namespace string) (daemonSet *appsv1.DaemonSet, err error) {
	daemonSet, err = client.AppsV1().DaemonSets(namespace).Get(context.TODO(), daemonSetName, metav1.GetOptions{})
	if err != nil {
		zap.L().Error("S-GetDaemonSetDetail 获取DaemonSet详情失败, ", zap.Error(err))
		return nil, errors.New("获取DaemonSet详情失败, " + err.Error())
	}

	return daemonSet, nil
}
