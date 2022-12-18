package service

import (
	"context"
	"errors"

	"go.uber.org/zap"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// 设置deployment副本数
func (d *deployment) ScaleDeployment(client *kubernetes.Clientset, deploymentName, namespace string, scaleNum int) (replica int32, err error) {
	// 获取autoscalingv1.Scale类型的对象(也是一种api资源类型)，能点出当前的副本数
	// k8s.io/api/autoscaling/v1
	scale, err := client.AppsV1().Deployments(namespace).GetScale(context.TODO(), deploymentName, metav1.GetOptions{})
	if err != nil {
		zap.L().Error("S-ScaleDeployment 获取Deployment副本数信息失败", zap.Error(err))
		return 0, errors.New("获取Deployment副本数信息失败, " + err.Error())
	}

	// 修改副本数
	scale.Spec.Replicas = int32(scaleNum)
	// 更新副本数，传入scale对象
	newScale, err := client.AppsV1().Deployments(namespace).UpdateScale(context.TODO(), deploymentName, scale, metav1.UpdateOptions{})
	if err != nil {
		zap.L().Error("S-ScaleDeployment 更新Deployment副本数信息失败", zap.Error(err))
		return 0, errors.New("更新Deployment副本数信息失败, " + err.Error())
	}

	return newScale.Spec.Replicas, nil
}
