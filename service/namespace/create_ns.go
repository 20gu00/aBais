package service

import (
	"context"
	"errors"

	"go.uber.org/zap"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type NsCreate struct {
	Name    string `json:"name"`
	Cluster string `json:"cluster"`
}

func (*namespace) CreatePvc(client *kubernetes.Clientset, data *NsCreate) (err error) {
	ns := &corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: data.Name,
		},
	}

	_, err = client.CoreV1().Namespaces().Create(context.TODO(), ns, metav1.CreateOptions{})
	if err != nil {
		zap.L().Error("C-CreateNs 创建ns失败, ", zap.Error(err))
		return errors.New("创建ns失败" + err.Error())
	}

	return nil
}
