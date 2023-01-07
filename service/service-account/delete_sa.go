package service

import (
	"context"
	"errors"

	"go.uber.org/zap"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// 删除job
func (j *sa) DeleteSa(client *kubernetes.Clientset, saName, namespace string) (err error) {
	err = client.CoreV1().ServiceAccounts(namespace).Delete(context.TODO(), saName, metav1.DeleteOptions{})
	if err != nil {
		zap.L().Error("S-DeleteSa 删除sa失败", zap.Error(err))
		return errors.New("S-DeleteSa 删除sa失败, " + err.Error())
	}

	return nil
}
