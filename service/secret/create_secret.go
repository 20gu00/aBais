package service

import (
	"context"
	"errors"

	"go.uber.org/zap"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type SecretCreate struct {
	Name      string            `json:"name"`
	Namespace string            `json:"namespace"`
	Cluster   string            `json:"cluster"`
	Data      map[string]string `json:"data"`
	Type      string            `json:"type"`
	//Label     map[string]string `json:"label"`
}

func (s *secret) CreateSecret(client *kubernetes.Clientset, data *SecretCreate) (err error) {
	dataTemp := make(map[string][]byte) //map
	for idx, val := range data.Data {
		dataTemp[idx] = []byte(val)
	}
	secret := &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      data.Name,
			Namespace: data.Namespace,
			//Labels:    data.Label,
		},
		Data: dataTemp,
		Type: corev1.SecretType(data.Type),
	}

	_, err = client.CoreV1().Secrets(data.Namespace).Create(context.TODO(), secret, metav1.CreateOptions{})
	if err != nil {
		zap.L().Error("S-CreateSecret 创建Secret失败, ", zap.Error(err))
		return errors.New("创建Secret失败" + err.Error())
	}

	return nil
}
