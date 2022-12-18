package cm

import (
	"context"
	"encoding/json"
	"errors"

	"go.uber.org/zap"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// 更新configmap
func (c *configMap) UpdateConfigMap(client *kubernetes.Clientset, namespace, content string) (err error) {
	var configMap = &corev1.ConfigMap{}

	err = json.Unmarshal([]byte(content), configMap)
	if err != nil {
		zap.L().Error("S-UpdateConfigMap 反序列化失败", zap.Error(err))
		return errors.New("反序列化失败, " + err.Error())
	}

	_, err = client.CoreV1().ConfigMaps(namespace).Update(context.TODO(), configMap, metav1.UpdateOptions{})
	if err != nil {
		zap.L().Error("S-UpdateConfigMap 更新ConfigMap失败", zap.Error(err))
		return errors.New("更新ConfigMap失败, " + err.Error())
	}
	return nil
}
