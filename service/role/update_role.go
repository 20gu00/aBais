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

func (j *role) UpdateRole(client *kubernetes.Clientset, namespace, content string) (err error) {
	var role = &rbacv1.Role{}
	err = json.Unmarshal([]byte(content), role)
	if err != nil {
		zap.L().Error("S-UpdateRole 反序列化失败", zap.Error(err))
		return errors.New("反序列化失败" + err.Error())
	}

	_, err = client.RbacV1().Roles(namespace).Update(context.TODO(), role, metav1.UpdateOptions{})
	if err != nil {
		zap.L().Error("S-UpdateJob 更新role失败", zap.Error(err))
		return errors.New("Role" + err.Error())
	}
	return nil
}
