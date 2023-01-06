package service

import (
	"context"
	"errors"

	"go.uber.org/zap"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// 删除job
func (r *role) DeleteRole(client *kubernetes.Clientset, roleName, namespace string) (err error) {
	err = client.RbacV1().Roles(namespace).Delete(context.TODO(), roleName, metav1.DeleteOptions{})
	if err != nil {
		zap.L().Error("S-DeleteRole 删除role失败", zap.Error(err))
		return errors.New("S-DeleteRole 删除role失败, " + err.Error())
	}

	return nil
}
