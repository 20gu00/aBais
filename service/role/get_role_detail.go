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
func (r *role) GetRoleDetail(client *kubernetes.Clientset, roleName, namespace string) (role *rbacv1.Role, err error) {
	role, err = client.RbacV1().Roles(namespace).Get(context.TODO(), roleName, metav1.GetOptions{})
	if err != nil {
		zap.L().Error("S-GetRoleDetail 获取role详情失败", zap.Error(err))
		return nil, errors.New("获取role详情失败" + err.Error())
	}

	role.ManagedFields = []metav1.ManagedFieldsEntry{}
	role.GetObjectKind().SetGroupVersionKind(schema.GroupVersionKind{
		Kind:    "Role",
		Version: "rbac.authorization.k8s.io/v1",
	})
	return role, nil
}
