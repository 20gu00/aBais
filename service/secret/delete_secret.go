package secret

import (
	"context"
	"errors"
	"go.uber.org/zap"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// 删除secret
func (s *secret) DeleteSecret(client *kubernetes.Clientset, secretName, namespace string) (err error) {
	err = client.CoreV1().Secrets(namespace).Delete(context.TODO(), secretName, metav1.DeleteOptions{})
	if err != nil {
		zap.L().Error("S-DeleteSecret 删除Secret失败", zap.Error(err))
		return errors.New("删除Secret失败, " + err.Error())
	}

	return nil
}
