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

// 获取pod详情
func (p *pod) GetPodDetail(client *kubernetes.Clientset, podName, namespace string) (pod *corev1.Pod, err error) {
	pod, err = client.CoreV1().Pods(namespace).Get(context.TODO(), podName, metav1.GetOptions{}) //Background
	if err != nil {
		zap.L().Error("S-GetPodDetail 获取Pod详情失败", zap.Error(err))
		return nil, errors.New("获取Pod详情失败: " + err.Error())
	}

	//可选 后端处理成yaml "sigs.k8s.io/yaml" yaml.Marshal(pod) 编码成yaml字符串 []byte fmt.Println(string(yaml.Marshal(pod)))

	//managedFields【ManagedFieldsEntry array】：这主要用于内部管理，用户通常不需要设置或理解此字段。
	pod.ManagedFields = []metav1.ManagedFieldsEntry{}
	pod.GetObjectKind().SetGroupVersionKind(schema.GroupVersionKind{
		Kind:    "Pod",
		Version: "v1",
	})
	return pod, nil
}
