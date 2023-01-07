package service

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
func (r *roleBinding) GetRoleBindingDetail(client *kubernetes.Clientset, roleBindingName, namespace string) (roleBinding *rbacv1.RoleBinding, err error) {
	roleBinding, err = client.RbacV1().RoleBindings(namespace).Get(context.TODO(), roleBindingName, metav1.GetOptions{})
	if err != nil {
		zap.L().Error("S-GetRoleBindingDetail 获取role binding详情失败", zap.Error(err))
		return nil, errors.New("获取role binding详情失败" + err.Error())
	}

	roleBinding.ManagedFields = []metav1.ManagedFieldsEntry{}
	roleBinding.GetObjectKind().SetGroupVersionKind(schema.GroupVersionKind{
		Kind:    "RoleBinding",
		Version: "rbac.authorization.k8s.io/v1",
	})
	return roleBinding, nil
}
