package clusterRoleBinding

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
func (r *clusterRoleBinding) GetClusterRoleBindingDetail(client *kubernetes.Clientset, clusterRoleBindingName string) (clusterRoleBinding *rbacv1.ClusterRoleBinding, err error) {
	clusterRoleBinding, err = client.RbacV1().ClusterRoleBindings().Get(context.TODO(), clusterRoleBindingName, metav1.GetOptions{})
	if err != nil {
		zap.L().Error("S-GetClusterRoleBindingDetail 获取cluster role binding详情失败", zap.Error(err))
		return nil, errors.New("获取cluster role binding详情失败" + err.Error())
	}

	clusterRoleBinding.ManagedFields = []metav1.ManagedFieldsEntry{}
	clusterRoleBinding.GetObjectKind().SetGroupVersionKind(schema.GroupVersionKind{
		Kind:    "ClusterRoleBinding",
		Version: "rbac.authorization.k8s.io/v1",
	})
	return clusterRoleBinding, nil
}
