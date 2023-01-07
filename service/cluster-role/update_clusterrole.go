package cluster_role

import (
	"context"
	"encoding/json"
	"errors"

	"go.uber.org/zap"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func (j *clusterRole) UpdateClusterRole(client *kubernetes.Clientset, content string) (err error) {
	var clusterRole = &rbacv1.ClusterRole{}
	err = json.Unmarshal([]byte(content), clusterRole)
	if err != nil {
		zap.L().Error("S-UpdateClusterRole 反序列化失败", zap.Error(err))
		return errors.New("反序列化失败" + err.Error())
	}

	_, err = client.RbacV1().ClusterRoles().Update(context.TODO(), clusterRole, metav1.UpdateOptions{})
	if err != nil {
		zap.L().Error("S-UpdateClusterRole 更新cluster role失败", zap.Error(err))
		return errors.New("ClusterRole" + err.Error())
	}
	return nil
}
