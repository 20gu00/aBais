package service

import (
	"context"
	"encoding/json"
	"errors"

	"go.uber.org/zap"
	batchv1beta1 "k8s.io/api/batch/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func (j *cronjob) UpdateCronJob(client *kubernetes.Clientset, namespace, content string) (err error) {
	var cronJob = &batchv1beta1.CronJob{}
	err = json.Unmarshal([]byte(content), cronJob)
	if err != nil {
		zap.L().Error("S-UpdateCronJob 反序列化失败", zap.Error(err))
		return errors.New("反序列化失败" + err.Error())
	}

	_, err = client.BatchV1beta1().CronJobs(namespace).Update(context.TODO(), cronJob, metav1.UpdateOptions{})
	if err != nil {
		zap.L().Error("S-UpdateJob 更新cronJob失败", zap.Error(err))
		return errors.New("更新cronJob失败" + err.Error())
	}
	return nil
}
