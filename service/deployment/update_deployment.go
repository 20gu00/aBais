package service

import (
	"context"
	"encoding/json"
	"errors"

	"go.uber.org/zap"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// 更新deployment
func (d *deployment) UpdateDeployment(client *kubernetes.Clientset, namespace, content string) (err error) {
	var deploy = &appsv1.Deployment{}

	// 反序列化yaml
	err = json.Unmarshal([]byte(content), deploy)
	if err != nil {
		zap.L().Error("S-UpdateDeployment 反序列化失败", zap.Error(err))
		return errors.New("反序列化失败, " + err.Error())
	}

	_, err = client.AppsV1().Deployments(namespace).Update(context.TODO(), deploy, metav1.UpdateOptions{})
	if err != nil {
		zap.L().Error("S-UpdateDeployment 更新Deployment失败", zap.Error(err))
		return errors.New("更新Deployment失败" + err.Error())
	}
	return nil
}
