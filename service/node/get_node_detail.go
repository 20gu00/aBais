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

// 获取node详情
func (n *node) GetNodeDetail(client *kubernetes.Clientset, nodeName string) (node *corev1.Node, err error) {
	node, err = client.CoreV1().Nodes().Get(context.TODO(), nodeName, metav1.GetOptions{})
	if err != nil {
		zap.L().Error("S-GetNodeDetail 获取Node详情失败", zap.Error(err))
		return nil, errors.New("获取Node详情失败, " + err.Error())
	}

	node.ManagedFields = []metav1.ManagedFieldsEntry{}
	node.GetObjectKind().SetGroupVersionKind(schema.GroupVersionKind{
		Kind:    "Node",
		Version: "v1",
	})
	return node, nil
}
