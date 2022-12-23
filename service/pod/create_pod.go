package service

import (
	"context"
	"errors"

	"go.uber.org/zap"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type PodCreateParam struct {
	Name          string            `json:"name"`
	Namespace     string            `json:"namespace"`
	Image         string            `json:"image"`
	Label         map[string]string `json:"label"`
	Cpu           string            `json:"cpu"`
	Memory        string            `json:"memory"`
	ContainerPort int32             `json:"container_port"`
	Cluster       string            `json:"cluster"`
}

func (p *pod) CreatePod(client *kubernetes.Clientset, data *PodCreateParam) (err error) {
	// 后端提供了ns的接口,前段只提供存在的ns选项
	pod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      data.Name,
			Namespace: data.Namespace,
			Labels:    data.Label,
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:  data.Name,
					Image: data.Image,
					Ports: []corev1.ContainerPort{
						{
							Name: "http",
							// 均可tcp...
							//Protocol:      corev1.ProtocolTCP,
							ContainerPort: data.ContainerPort,
						},
					},
					Resources: corev1.ResourceRequirements{
						Limits: map[corev1.ResourceName]resource.Quantity{
							corev1.ResourceCPU:    resource.MustParse(data.Cpu),
							corev1.ResourceMemory: resource.MustParse(data.Memory),
						},
						Requests: map[corev1.ResourceName]resource.Quantity{
							corev1.ResourceCPU:    resource.MustParse(data.Cpu),
							corev1.ResourceMemory: resource.MustParse(data.Memory),
						},
					},
				},
			},
		},
		Status: corev1.PodStatus{},
	}
	_, err = client.CoreV1().Pods(data.Namespace).Create(context.TODO(), pod, metav1.CreateOptions{})
	if err != nil {
		zap.L().Error("C-CreatePod CreatePod, ", zap.Error(err))
		return errors.New("CreatePod, " + err.Error())
	}

	return nil
}
