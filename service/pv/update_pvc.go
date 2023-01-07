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
func (p *pv) UpdatePv(client *kubernetes.Clientset, content string) (err error) {
	var pv = &corev1.PersistentVolume{}
	err = json.Unmarshal([]byte(content), pv)
	if err != nil {
		zap.L().Error("S-UpdatePv 反序列化失败", zap.Error(err))
		return errors.New("反序列化失败, " + err.Error())
	}

	_, err = client.CoreV1().PersistentVolumes().Update(context.TODO(), pv, metav1.UpdateOptions{})
	if err != nil {
		zap.L().Error("S-UpdatePv 更新Pv失败", zap.Error(err))
		return errors.New("更新Pv失败, " + err.Error())
	}
	return nil
}
