package statefulset

import (
	"context"
	"errors"

	"go.uber.org/zap"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// 删除statefulset
func (s *statefulSet) DeleteStatefulSet(client *kubernetes.Clientset, statefulSetName, namespace string) (err error) {
	err = client.AppsV1().StatefulSets(namespace).Delete(context.TODO(), statefulSetName, metav1.DeleteOptions{})
	if err != nil {
		zap.L().Error("S-DeleteStatefulSet 删除StatefulSet失败", zap.Error(err))
		return errors.New("删除StatefulSet失败, " + err.Error())
	}

	return nil
}
