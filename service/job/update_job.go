package service

import (
	"context"
	"encoding/json"
	"errors"

	"go.uber.org/zap"
	batchv1 "k8s.io/api/batch/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func (j *job) UpdateJob(client *kubernetes.Clientset, namespace, content string) (err error) {
	var job = &batchv1.Job{}
	err = json.Unmarshal([]byte(content), job)
	if err != nil {
		zap.L().Error("S-UpdateJob 反序列化失败", zap.Error(err))
		return errors.New("反序列化失败" + err.Error())
	}

	_, err = client.BatchV1().Jobs(namespace).Update(context.TODO(), job, metav1.UpdateOptions{})
	if err != nil {
		zap.L().Error("S-UpdateJob 更新Job失败", zap.Error(err))
		return errors.New("更新Job失败" + err.Error())
	}
	return nil
}
