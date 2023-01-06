package service

import (
	"context"
	"errors"
	"strings"

	"go.uber.org/zap"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type RoleCreate struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	Cluster   string `json:"cluster"`
	ApiGroup  string `json:"api_group"`
	Resources string `json:"resources"`
	verbs     string `json:"verbs"`
}

func (s *role) CreateRole(client *kubernetes.Clientset, data *RoleCreate) (err error) {
	apiGroups := strings.Split(data.ApiGroup, "|")
	resources := strings.Split(data.Resources, "|")
	verbs := strings.Split(data.verbs, "|")
	role := &rbacv1.Role{
		ObjectMeta: metav1.ObjectMeta{
			Name:      data.Name,
			Namespace: data.Namespace,
		},
		Rules: []rbacv1.PolicyRule{
			{
				APIGroups: apiGroups,
				Verbs:     resources,
				Resources: verbs,
			},
		},
	}

	_, err = client.RbacV1().Roles(data.Namespace).Create(context.TODO(), role, metav1.CreateOptions{})
	if err != nil {
		zap.L().Error("S-CreateRole 创建Role失败, ", zap.Error(err))
		return errors.New("创建Role失败" + err.Error())
	}

	return nil
}
