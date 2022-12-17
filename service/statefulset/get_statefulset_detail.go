package statefulset

import (
	"context"
	"errors"

	"go.uber.org/zap"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// 获取statefulset详情
func (s *statefulSet) GetStatefulSetDetail(client *kubernetes.Clientset, statefulSetName, namespace string) (statefulSet *appsv1.StatefulSet, err error) {
	statefulSet, err = client.AppsV1().StatefulSets(namespace).Get(context.TODO(), statefulSetName, metav1.GetOptions{})
	if err != nil {
		zap.L().Error("S-GetStatefulSetDetail 获取StatefulSet详情失败, ", zap.Error(err))
		return nil, errors.New("获取StatefulSet详情失败, " + err.Error())
	}

	return statefulSet, nil
}
