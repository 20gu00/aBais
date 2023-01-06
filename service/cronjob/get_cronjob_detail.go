package service

import (
	"context"
	"errors"

	"go.uber.org/zap"
	batchv1beta1 "k8s.io/api/batch/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/kubernetes"
)

// yaml
func (j *cronjob) GetCronJobDetail(client *kubernetes.Clientset, cronJobName, namespace string) (cronJob *batchv1beta1.CronJob, err error) {
	cronJob, err = client.BatchV1beta1().CronJobs(namespace).Get(context.TODO(), cronJobName, metav1.GetOptions{})
	if err != nil {
		zap.L().Error("S-GetCronJobDetail 获取cronjob详情失败", zap.Error(err))
		return nil, errors.New("获取cronjob详情失败" + err.Error())
	}

	cronJob.ManagedFields = []metav1.ManagedFieldsEntry{}
	cronJob.GetObjectKind().SetGroupVersionKind(schema.GroupVersionKind{
		Kind:    "Job",
		Version: "batch/v1beta1",
	})
	return cronJob, nil
}
