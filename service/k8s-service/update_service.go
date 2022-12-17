package k8sSvc

import (
	"context"
	"encoding/json"
	"errors"

	"go.uber.org/zap"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// 更新service
func (s *k8sService) UpdateService(client *kubernetes.Clientset, namespace, content string) (err error) {
	var service = &corev1.Service{}
	err = json.Unmarshal([]byte(content), service)
	if err != nil {
		zap.L().Error("反序列化失败", zap.Error(err))
		return errors.New("反序列化失败, " + err.Error())
	}

	_, err = client.CoreV1().Services(namespace).Update(context.TODO(), service, metav1.UpdateOptions{})
	if err != nil {
		zap.L().Error("更新service失败", zap.Error(err))
		return errors.New("更新service失败, " + err.Error())
	}
	return nil
}
