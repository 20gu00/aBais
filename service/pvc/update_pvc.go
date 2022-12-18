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

// 更新pvc
func (p *pvc) UpdatePvc(client *kubernetes.Clientset, namespace, content string) (err error) {
	var pvc = &corev1.PersistentVolumeClaim{}
	err = json.Unmarshal([]byte(content), pvc)
	if err != nil {
		zap.L().Error("S-UpdatePvc 反序列化失败", zap.Error(err))
		return errors.New("反序列化失败, " + err.Error())
	}

	_, err = client.CoreV1().PersistentVolumeClaims(namespace).Update(context.TODO(), pvc, metav1.UpdateOptions{})
	if err != nil {
		zap.L().Error("S-UpdatePvc 更新Pvc失败", zap.Error(err))
		return errors.New("更新Pvc失败, " + err.Error())
	}
	return nil
}
