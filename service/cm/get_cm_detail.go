package service

import (
	"context"
	"errors"
	"k8s.io/apimachinery/pkg/runtime/schema"

	"go.uber.org/zap"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// 获取configmap详情
func (c *configMap) GetConfigMapDetail(client *kubernetes.Clientset, configMapName, namespace string) (configMap *corev1.ConfigMap, err error) {
	configMap, err = client.CoreV1().ConfigMaps(namespace).Get(context.TODO(), configMapName, metav1.GetOptions{})
	if err != nil {
		zap.L().Error("S-GetConfigMapDetail 获取ConfigMap详情失败", zap.Error(err))
		return nil, errors.New("获取ConfigMap详情失败" + err.Error())
	}

	configMap.ManagedFields = []metav1.ManagedFieldsEntry{}
	configMap.GetObjectKind().SetGroupVersionKind(schema.GroupVersionKind{
		Kind:    "ConfigMap",
		Version: "apps/v1",
	})
	return configMap, nil
}
