package service
import (
	"context"
	"encoding/json"
	"errors"

	"go.uber.org/zap"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func (j *roleBinding) UpdateRoleBinding(client *kubernetes.Clientset, namespace, content string) (err error) {
	var roleBinding = &rbacv1.RoleBinding{}
	err = json.Unmarshal([]byte(content), roleBinding)
	if err != nil {
		zap.L().Error("S-UpdateRoleBinding 反序列化失败", zap.Error(err))
		return errors.New("反序列化失败" + err.Error())
	}

	_, err = client.RbacV1().RoleBindings(namespace).Update(context.TODO(), roleBinding, metav1.UpdateOptions{})
	if err != nil {
		zap.L().Error("S-UpdateRoleBinding 更新role binding失败", zap.Error(err))
		return errors.New("更新role binding失败" + err.Error())
	}
	return nil
}