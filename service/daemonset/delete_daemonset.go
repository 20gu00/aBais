package daemonset

import (
	"context"
	"errors"

	"go.uber.org/zap"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// 删除daemonset
func (d *daemonSet) DeleteDaemonSet(client *kubernetes.Clientset, daemonSetName, namespace string) (err error) {
	err = client.AppsV1().DaemonSets(namespace).Delete(context.TODO(), daemonSetName, metav1.DeleteOptions{})
	if err != nil {
		zap.L().Error("S-DeleteDaemonSet 删除DaemonSet失败, ", zap.Error(err))
		return errors.New("删除DaemonSet失败, " + err.Error())
	}

	return nil
}
