package service

import (
	"context"
	"errors"

	"go.uber.org/zap"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// 删除job
func (r *roleBinding) DeleteRoleBinding(client *kubernetes.Clientset, roleBindingName, namespace string) (err error) {
	err = client.RbacV1().Roles(namespace).Delete(context.TODO(), roleBindingName, metav1.DeleteOptions{})
	if err != nil {
		zap.L().Error("S-DeleteRoleBinding 删除role binding失败", zap.Error(err))
		return errors.New("S-DeleteRoleBinding 删除role binding失败, " + err.Error())
	}

	return nil
}
