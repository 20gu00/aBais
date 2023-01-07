package clusterRoleBinding

import (
	"context"
	"errors"
	"go.uber.org/zap"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type ClusterRoleBindingCreate struct {
	Name            string `json:"name"`
	Cluster         string `json:"cluster"`
	ClusterRoleName string `json:"clusterrole_name"`
	SaName          string `json:"sa_name"`
	SaNamespace     string `json:"sa_namespace"`
}

func (s *clusterRoleBinding) CreateClusterRoleBinding(client *kubernetes.Clientset, data *ClusterRoleBindingCreate) (err error) {
	clusterRoleBinding := &rbacv1.ClusterRoleBinding{
		ObjectMeta: metav1.ObjectMeta{
			Name: data.Name,
		},
		Subjects: []rbacv1.Subject{
			{
				Kind:      rbacv1.ServiceAccountKind,
				Name:      data.SaName,
				Namespace: data.SaNamespace,
			},
		},
		RoleRef: rbacv1.RoleRef{
			Kind:     "ClusterRole",
			Name:     data.ClusterRoleName,
			APIGroup: "rbac.authorization.k8s.io",
		},
	}

	_, err = client.RbacV1().ClusterRoleBindings().Create(context.TODO(), clusterRoleBinding, metav1.CreateOptions{})
	if err != nil {
		zap.L().Error("S-CreateClusterRoleBinding 创建Cluster Role Binding失败, ", zap.Error(err))
		return errors.New("创建Cluster Role Binding失败" + err.Error())
	}

	return nil
}
