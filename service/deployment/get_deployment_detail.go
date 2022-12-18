package service

import (
	"context"
	"errors"
	"go.uber.org/zap"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

//获取deployment详情
func (d *deployment) GetDeploymentDetail(client *kubernetes.Clientset, deploymentName, namespace string) (deployment *appsv1.Deployment, err error) {
	deployment, err = client.AppsV1().Deployments(namespace).Get(context.TODO(), deploymentName, metav1.GetOptions{})
	if err != nil {
		zap.L().Error("S-GetDeploymentDetail 获取Deployment详情失败", zap.Error(err))
		return nil, errors.New("获取Deployment详情失败, " + err.Error())
	}

	return deployment, nil
}
