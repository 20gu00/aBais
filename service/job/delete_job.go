package service

import (
	"context"
	"errors"

	"go.uber.org/zap"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// 删除job
func (j *job) DeleteJob(client *kubernetes.Clientset, jobName, namespace string) (err error) {
	err = client.BatchV1().Jobs(namespace).Delete(context.TODO(), jobName, metav1.DeleteOptions{})
	if err != nil {
		zap.L().Error("S-DeleteJob 删除Job失败", zap.Error(err))
		return errors.New("S-DeleteJob 删除Job失败, " + err.Error())
	}

	return nil
}
