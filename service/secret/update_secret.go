package secret

import (
	"context"
	"encoding/json"
	"errors"
	"go.uber.org/zap"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// 更新secret
func (s *secret) UpdateSecret(client *kubernetes.Clientset, namespace, content string) (err error) {
	var secret = &corev1.Secret{}

	err = json.Unmarshal([]byte(content), secret)
	if err != nil {
		zap.L().Error("S-UpdateSecret 反序列化失败", zap.Error(err)
		err.Error()))
		return errors.New("反序列化失败, " + err.Error())
	}

	_, err = client.CoreV1().Secrets(namespace).Update(context.TODO(), secret, metav1.UpdateOptions{})
	if err != nil {
		zap.L().Error("S-UpdateSecret 更新Secret失败, ", zap.Error(err))
		return errors.New("更新Secret失败, " + err.Error())
	}
	return nil
}
