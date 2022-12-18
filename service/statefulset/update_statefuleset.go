package service

import (
	"context"
	"encoding/json"
	"errors"

	"go.uber.org/zap"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// 更新statefulset
func (s *statefulSet) UpdateStatefulSet(client *kubernetes.Clientset, namespace, content string) (err error) {
	var statefulSet = &appsv1.StatefulSet{}

	err = json.Unmarshal([]byte(content), statefulSet)
	if err != nil {
		zap.L().Error("C-UpdateStatefulSet 反序列化失败", zap.Error(err))
		return errors.New("反序列化失败, " + err.Error())
	}

	_, err = client.AppsV1().StatefulSets(namespace).Update(context.TODO(), statefulSet, metav1.UpdateOptions{})
	if err != nil {
		zap.L().Error("C-UpdateStatefulSet 更新StatefulSet失败, ", zap.Error(err))
		return errors.New("更新StatefulSet失败, " + err.Error())
	}
	return nil
}
