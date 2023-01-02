package service

import (
	"context"
	"errors"
	"go.uber.org/zap"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type ConfigmapCreate struct {
	Name      string            `json:"name"`
	Namespace string            `json:"namespace"`
	Cluster   string            `json:"cluster"`
	Data      map[string]string `json:"data"`
}

func (d *configMap) CreateCm(client *kubernetes.Clientset, data *configmapCreate) (err error) {
	cm := &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      data.Name,
			Namespace: data.Namespace,
		},
		Data: data.Data,
	}

	_, err = client.CoreV1().ConfigMaps(data.Namespace).Create(context.TODO(), cm, metav1.CreateOptions{})
	if err != nil {
		zap.L().Error("C-CreateCm 创建ConfigMap失败, ", zap.Error(err))
		return errors.New("创建ConfigMap失败" + err.Error())
	}

	return nil
}
