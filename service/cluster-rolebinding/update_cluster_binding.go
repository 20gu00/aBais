package clusterRoleBinding

import (
	"context"
	"encoding/json"
	"errors"

	"go.uber.org/zap"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func (j *clusterRoleBinding) UpdateClusteroleBinding(client *kubernetes.Clientset, content string) (err error) {
	var clusterRoleBinding = &rbacv1.ClusterRoleBinding{}
	err = json.Unmarshal([]byte(content), clusterRoleBinding)
	if err != nil {
		zap.L().Error("S-UpdateClusteroleBinding 反序列化失败", zap.Error(err))
		return errors.New("反序列化失败" + err.Error())
	}

	_, err = client.RbacV1().ClusterRoleBindings().Update(context.TODO(), clusterRoleBinding, metav1.UpdateOptions{})
	if err != nil {
		zap.L().Error("S-UpdateClusteroleBinding 更新cluster role binding失败", zap.Error(err))
		return errors.New("更新cluster role binding失败" + err.Error())
	}
	return nil
}
