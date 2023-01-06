package service

import (
	"context"
	"errors"

	"go.uber.org/zap"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// 删除job
func (j *cronjob) DeleteCronJob(client *kubernetes.Clientset, cronJobName, namespace string) (err error) {
	err = client.BatchV1beta1().CronJobs(namespace).Delete(context.TODO(), cronJobName, metav1.DeleteOptions{})
	if err != nil {
		zap.L().Error("S-DeleteCronJob 删除cronJob失败", zap.Error(err))
		return errors.New("S-DeleteCronJob 删除cronJob失败, " + err.Error())
	}

	return nil
}
