package service

import (
	"context"
	"errors"

	"go.uber.org/zap"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type SaCreate struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	Cluster   string `json:"cluster"`
}

func (c *sa) CreateSa(client *kubernetes.Clientset, data *SaCreate) (err error) {
	sa := &corev1.ServiceAccount{
		ObjectMeta: metav1.ObjectMeta{
			Name:      data.Name,
			Namespace: data.Namespace,
		},
	}

	_, err = client.CoreV1().ServiceAccounts(data.Namespace).Create(context.TODO(), sa, metav1.CreateOptions{})
	if err != nil {
		zap.L().Error("S-CreateSa 创建sa失败, ", zap.Error(err))
		return errors.New("创建sa失败" + err.Error())
	}

	return nil
}
