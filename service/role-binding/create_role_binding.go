package service

import (
	"context"
	"errors"
	"go.uber.org/zap"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type RoleBindingCreate struct {
	Name        string `json:"name"`
	Namespace   string `json:"namespace"`
	Cluster     string `json:"cluster"`
	RoleName    string `json:"role_name"`
	SaName      string `json:"sa_name"`
	SaNamespace string `json:"sa_namespace"`
}

func (s *roleBinding) CreateRoleBinding(client *kubernetes.Clientset, data *RoleBindingCreate) (err error) {
	roleBinding := &rbacv1.RoleBinding{
		ObjectMeta: metav1.ObjectMeta{
			Name:      data.Name,
			Namespace: data.Namespace,
		},
		Subjects: []rbacv1.Subject{
			{
				Kind:      rbacv1.ServiceAccountKind,
				Name:      data.SaName,
				Namespace: data.SaNamespace,
			},
		},
		RoleRef: rbacv1.RoleRef{
			Kind:     "Role",
			Name:     data.RoleName,
			APIGroup: "rbac.authorization.k8s.io",
		},
	}

	_, err = client.RbacV1().RoleBindings(data.Namespace).Create(context.TODO(), roleBinding, metav1.CreateOptions{})
	if err != nil {
		zap.L().Error("S-CreateRoleBinding 创建Role Binding失败, ", zap.Error(err))
		return errors.New("创建Role Binding失败" + err.Error())
	}

	return nil
}
