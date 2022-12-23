package service

import (
	"context"
	"errors"
	"go.uber.org/zap"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"strconv"
	"strings"
)

type PodCreateParam struct {
	Name          string            `json:"name"`
	Namespace     string            `json:"namespace"`
	Image         string            `json:"image"`
	Label         map[string]string `json:"label"`
	LimitCpu      string            `json:"limit_cpu"`
	LimitMemory   string            `json:"limit_memory"`
	RequestCpu    string            `json:"request_cpu"`
	RequestMemory string            `json:"request_memory"`
	ContainerPort string            `json:"container_port"`
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
			Containers: CreateContainer(data),
		},
		Status: corev1.PodStatus{},
	}
	_, err = client.CoreV1().Pods(data.Namespace).Create(context.TODO(), pod, metav1.CreateOptions{})
	if err != nil {
		zap.L().Error("C-CreatePod 创建pod失败", zap.Error(err))
		return errors.New("创建pod失败, " + err.Error())
	}

	return nil
}

func CreateContainer(data *PodCreateParam) []corev1.Container {
	// image limit request对应
	var (
		imgs           = strings.Split(data.Image, ",")
		limitCpus      = strings.Split(data.LimitCpu, ",")
		limitMems      = strings.Split(data.LimitMemory, ",")
		ReqCpus        = strings.Split(data.RequestCpu, ",")
		ReqMems        = strings.Split(data.RequestMemory, ",")
		containerPorts = strings.Split(data.ContainerPort, ",")
		ports          = []int32{}
	)

	for _, item := range containerPorts {
		v, _ := strconv.Atoi(item)
		ports = append(ports, int32(v))
	}

	containers := []corev1.Container{}

	for i, _ := range imgs {
		c := corev1.Container{
			Name:  data.Name + "-" + imgs[i],
			Image: imgs[i],
			Ports: []corev1.ContainerPort{
				{
					// 容器端口和容器这里是一对一
					Name: imgs[i] + "port1", // "http",
					// 均可tcp...
					// Protocol:      corev1.ProtocolTCP,
					// containerPort--targetPort(pod)
					ContainerPort: ports[i],
				},
			},
			Resources: corev1.ResourceRequirements{
				Limits: map[corev1.ResourceName]resource.Quantity{
					corev1.ResourceCPU:    resource.MustParse(limitCpus[i]),
					corev1.ResourceMemory: resource.MustParse(limitMems[i]),
				},
				Requests: map[corev1.ResourceName]resource.Quantity{
					corev1.ResourceCPU:    resource.MustParse(ReqCpus[i]), // 1Gi
					corev1.ResourceMemory: resource.MustParse(ReqMems[i]),
				},
			},
		}
		containers = append(containers, c)
	}

	return containers
}
