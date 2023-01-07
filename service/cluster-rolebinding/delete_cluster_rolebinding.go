package clusterRoleBinding

import (
	"context"
	"errors"

	"go.uber.org/zap"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// 删除job
func (r *clusterRoleBinding) DeleteClusterRoleBinding(client *kubernetes.Clientset, clusterRoleBindingName string) (err error) {
	err = client.RbacV1().ClusterRoleBindings().Delete(context.TODO(), clusterRoleBindingName, metav1.DeleteOptions{})
	if err != nil {
		zap.L().Error("S-DeleteClusterRoleBinding 删除cluster role binding失败", zap.Error(err))
		return errors.New("S-DeleteClusterRoleBinding 删除cluster role binding失败, " + err.Error())
	}

	return nil
}
