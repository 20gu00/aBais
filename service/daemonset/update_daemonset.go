package daemonset

import (
	"context"
	"encoding/json"
	"errors"

	"go.uber.org/zap"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// 更新daemonset
func (d *daemonSet) UpdateDaemonSet(client *kubernetes.Clientset, namespace, content string) (err error) {
	var daemonSet = &appsv1.DaemonSet{}

	err = json.Unmarshal([]byte(content), daemonSet)
	if err != nil {
		zap.L().Error("S-UpdateDaemonSet 反序列化失败, ", zap.Error(err))
		return errors.New("反序列化失败, " + err.Error())
	}

	_, err = client.AppsV1().DaemonSets(namespace).Update(context.TODO(), daemonSet, metav1.UpdateOptions{})
	if err != nil {
		zap.L().Error("S-UpdateDaemonSet 更新DaemonSet失败, ", zap.Error(err))
		return errors.New("更新DaemonSet失败, " + err.Error())
	}
	return nil
}
