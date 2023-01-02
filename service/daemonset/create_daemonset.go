package service

import (
	"context"
	"errors"
	"strconv"
	"strings"

	"go.uber.org/zap"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/kubernetes"
)

type DaemonsetCreate struct {
	Name          string            `json:"name"`
	Namespace     string            `json:"namespace"`
	Image         string            `json:"image"`
	Label         map[string]string `json:"label"`
	PodLabel      map[string]string `json:"pod_label"`
	LimitCpu      string            `json:"limit_cpu"`
	LimitMemory   string            `json:"limit_memory"`
	RequestCpu    string            `json:"request_cpu"`
	RequestMemory string            `json:"request_memory"`
	ContainerPort string            `json:"container_port"`
	HealthCheck   bool              `json:"health_check"`
	HealthPath    string            `json:"health_path"`
	Cluster       string            `json:"cluster"`
}

func (d *daemonSet) CreateDaemonset(client *kubernetes.Clientset, data *DaemonsetCreate) (err error) {
	daemonset := &appsv1.DaemonSet{
		ObjectMeta: metav1.ObjectMeta{
			Name:      data.Name,
			Namespace: data.Namespace,
			Labels:    data.Label,
		},

		Spec: appsv1.DaemonSetSpec{
			Selector: &metav1.LabelSelector{
				MatchLabels: data.PodLabel,
			},
			// pod template
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Name:   data.Name,
					Labels: data.PodLabel,
				},

				Spec: corev1.PodSpec{
					Containers: CreateContainer(data),
				},
			},
		},
		Status: appsv1.DaemonSetStatus{},
	}

	_, err = client.AppsV1().DaemonSets(data.Namespace).Create(context.TODO(), daemonset, metav1.CreateOptions{})
	if err != nil {
		zap.L().Error("C-CreateDaemonset 创建Daemonset失败, ", zap.Error(err))
		return errors.New("创建Daemonset失败" + err.Error())
	}

	return nil
}

func CreateContainer(data *DaemonsetCreate) []corev1.Container {
	var (
		imgs           = strings.Split(data.Image, ",")
		LimitCpus      = strings.Split(data.LimitCpu, ",")
		LimitMems      = strings.Split(data.LimitMemory, ",")
		ReqCpus        = strings.Split(data.RequestCpu, ",")
		ReqMems        = strings.Split(data.RequestMemory, ",")
		containerPorts = strings.Split(data.ContainerPort, ",")
		ports          = []int32{}
	)

	//对空字符即使len为0,切割  得到[]长度为1
	//空字符串以比如,切割得到的切片长度为1
	//fmt.Println(limitCpus, data.LimitCpu, "1", len(data.LimitCpu), len(limitCpus))
	for _, item := range containerPorts {
		v, _ := strconv.Atoi(item)
		ports = append(ports, int32(v))
	}

	containers := []corev1.Container{}
	for i, _ := range imgs {
		c := corev1.Container{
			Name:  data.Name + imgs[i] + "-" + strconv.Itoa(i),
			Image: imgs[i],
		}

		if i < len(ports) && data.ContainerPort != "" {
			c.Ports = []corev1.ContainerPort{
				{
					Name:          imgs[i] + "-port-1",
					Protocol:      corev1.ProtocolTCP,
					ContainerPort: ports[i],
				},
			}
		}

		if i < len(LimitMems) && data.LimitMemory != "" {
			c.Resources = corev1.ResourceRequirements{
				Limits: map[corev1.ResourceName]resource.Quantity{corev1.ResourceMemory: resource.MustParse(LimitMems[i])},
			}
		}

		if i < len(LimitCpus) && data.LimitCpu != "" {
			c.Resources = corev1.ResourceRequirements{
				Limits: map[corev1.ResourceName]resource.Quantity{
					corev1.ResourceCPU: resource.MustParse(LimitCpus[i]),
				},
			}
		}

		if i < len(ReqMems) && data.RequestMemory != "" {
			c.Resources = corev1.ResourceRequirements{
				Requests: map[corev1.ResourceName]resource.Quantity{
					corev1.ResourceMemory: resource.MustParse(ReqMems[i]),
				},
			}
		}

		if i < len(ReqCpus) && data.RequestCpu != "" {
			c.Resources = corev1.ResourceRequirements{
				Requests: map[corev1.ResourceName]resource.Quantity{
					corev1.ResourceCPU: resource.MustParse(ReqCpus[i]),
				},
			}
		}
		if data.HealthCheck {
			c.ReadinessProbe = &corev1.Probe{
				ProbeHandler: corev1.ProbeHandler{
					HTTPGet: &corev1.HTTPGetAction{
						Path: data.HealthPath,
						Port: intstr.IntOrString{
							Type:   0,
							IntVal: ports[i],
						},
					},
				},
				InitialDelaySeconds: 5,
				TimeoutSeconds:      5,
				PeriodSeconds:       5,
			}
			c.LivenessProbe = &corev1.Probe{
				ProbeHandler: corev1.ProbeHandler{
					HTTPGet: &corev1.HTTPGetAction{
						Path: data.HealthPath,
						Port: intstr.IntOrString{
							Type:   0,
							IntVal: ports[i],
						},
					},
				},
				InitialDelaySeconds: 15,
				TimeoutSeconds:      5,
				PeriodSeconds:       5,
			}
		}
		containers = append(containers, c)
	}
	return containers
}
