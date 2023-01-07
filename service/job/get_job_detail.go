package service

import (
	"context"
	"errors"

	"go.uber.org/zap"
	batchv1 "k8s.io/api/batch/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/kubernetes"
)

// yaml
func (j *job) GetJobDetail(client *kubernetes.Clientset, JobName, namespace string) (job *batchv1.Job, err error) {
	job, err = client.BatchV1().Jobs(namespace).Get(context.TODO(), JobName, metav1.GetOptions{})
	if err != nil {
		zap.L().Error("S-GetJobDetail 获取job详情失败", zap.Error(err))
		return nil, errors.New("获取job详情失败" + err.Error())
	}

	job.ManagedFields = []metav1.ManagedFieldsEntry{}
	job.GetObjectKind().SetGroupVersionKind(schema.GroupVersionKind{
		Kind:    "Job",
		Version: "batch/v1",
	})
	return job, nil
}
