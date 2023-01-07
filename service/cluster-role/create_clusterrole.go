package cluster_role

import (
	"context"
	"errors"
	"strings"

	"go.uber.org/zap"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type ClusterRoleCreate struct {
	Name      string `json:"name"`
	Cluster   string `json:"cluster"`
	ApiGroup  string `json:"api_group"`
	Resources string `json:"resources"`
	Verbs     string `json:"verbs"`
}

func (s *clusterRole) CreateClusterRole(client *kubernetes.Clientset, data *ClusterRoleCreate) (err error) {
	apiGroups := strings.Split(data.ApiGroup, "|")
	resources := strings.Split(data.Resources, "|")
	verbs := strings.Split(data.Verbs, "|")
	clusterRole := &rbacv1.ClusterRole{
		ObjectMeta: metav1.ObjectMeta{
			Name: data.Name,
		},
		Rules: []rbacv1.PolicyRule{
			{
				APIGroups: apiGroups,
				Verbs:     resources,
				Resources: verbs,
			},
		},
	}

	_, err = client.RbacV1().ClusterRoles().Create(context.TODO(), clusterRole, metav1.CreateOptions{})
	if err != nil {
		zap.L().Error("S-CreateClusterRole 创建ClusterRole失败, ", zap.Error(err))
		return errors.New("创建clusterRole失败" + err.Error())
	}

	return nil
}
