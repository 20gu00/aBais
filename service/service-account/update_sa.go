package service

import (
	"context"
	"encoding/json"
	"errors"

	"go.uber.org/zap"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func (j *sa) UpdateSa(client *kubernetes.Clientset, namespace, content string) (err error) {
	var sa = &corev1.ServiceAccount{}
	err = json.Unmarshal([]byte(content), sa)
	if err != nil {
		zap.L().Error("S-UpdateSa 反序列化失败", zap.Error(err))
		return errors.New("反序列化失败" + err.Error())
	}

	_, err = client.CoreV1().ServiceAccounts(namespace).Update(context.TODO(), sa, metav1.UpdateOptions{})
	if err != nil {
		zap.L().Error("S-UpdateSa 更新sa失败", zap.Error(err))
		return errors.New("更新sa失败" + err.Error())
	}
	return nil
}
