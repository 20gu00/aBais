package cm

import (
	"context"
	"errors"

	"go.uber.org/zap"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// 删除configmap
func (c *configMap) DeleteConfigMap(client *kubernetes.Clientset, configMapName, namespace string) (err error) {
	err = client.CoreV1().ConfigMaps(namespace).Delete(context.TODO(), configMapName, metav1.DeleteOptions{})
	if err != nil {
		zap.L().Error("S-DeleteConfigMap 删除ConfigMap失败", zap.Error(err))
		return errors.New("删除ConfigMap失败, " + err.Error())
	}

	return nil
}
