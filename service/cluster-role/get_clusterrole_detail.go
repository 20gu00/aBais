package cluster_role

import (
	"context"
	"errors"

	"go.uber.org/zap"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/kubernetes"
)

// yaml
func (r *clusterRole) GetClusterRoleDetail(client *kubernetes.Clientset, clusterRoleName string) (clusterRole *rbacv1.ClusterRole, err error) {
	clusterRole, err = client.RbacV1().ClusterRoles().Get(context.TODO(), clusterRoleName, metav1.GetOptions{})
	if err != nil {
		zap.L().Error("S-GetClusterRoleDetail 获取clusterRole详情失败", zap.Error(err))
		return nil, errors.New("获取clusterRole详情失败" + err.Error())
	}

	clusterRole.ManagedFields = []metav1.ManagedFieldsEntry{}
	clusterRole.GetObjectKind().SetGroupVersionKind(schema.GroupVersionKind{
		Kind:    "ClusterRole",
		Version: "rbac.authorization.k8s.io/v1",
	})
	return clusterRole, nil
}
