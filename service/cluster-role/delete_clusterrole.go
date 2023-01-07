package cluster_role

import (
	"context"
	"errors"

	"go.uber.org/zap"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// 删除job
func (r *clusterRole) DeleteClusterRole(client *kubernetes.Clientset, clusterRoleName string) (err error) {
	err = client.RbacV1().ClusterRoles().Delete(context.TODO(), clusterRoleName, metav1.DeleteOptions{})
	if err != nil {
		zap.L().Error("S-DeleteClusterRole 删除clusterRole失败", zap.Error(err))
		return errors.New("S-DeleteClusterRole 删除clusterRole失败, " + err.Error())
	}

	return nil
}
