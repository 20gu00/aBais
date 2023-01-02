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

type StatefulsetCreate struct {
	Name            string            `json:"name"`
	Namespace       string            `json:"namespace"`
	Image           string            `json:"image"`
	Label           map[string]string `json:"label"`
	PodLabel        map[string]string `json:"pod_label"`
	LimitCpu        string            `json:"limit_cpu"`
	LimitMemory     string            `json:"limit_memory"`
	RequestCpu      string            `json:"request_cpu"`
	RequestMemory   string            `json:"request_memory"`
	ContainerPort   string            `json:"container_port"`
	HealthCheck     bool              `json:"health_check"`
	HealthPath      string            `json:"health_path"`
	Cluster         string            `json:"cluster"`
	ServiceName     string            `json:"service_name"`
	Replicas        int32             `json:"replicas"`
	MountName       string            `json:"volume_mount_name"`
	MountPath       string            `json:"mount_path"`
	VolumeClaimName string            `json:"volume_claim_name"`
	AccessMode      string            `json:"access_mode"`
	Storage         string            `json:"storage"`
}

func (d *statefulSet) CreateDaemonset(client *kubernetes.Clientset, data *StatefulsetCreate) (err error) {
	statefulset := &appsv1.StatefulSet{
		ObjectMeta: metav1.ObjectMeta{
			Name:      data.Name,
			Namespace: data.Namespace,
			Labels:    data.Label,
		},

		Spec: appsv1.StatefulSetSpec{
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
			VolumeClaimTemplates: CreateVolumeClaim(data),
		},
		Status: appsv1.StatefulSetStatus{},
	}

	_, err = client.AppsV1().StatefulSets(data.Namespace).Create(context.TODO(), statefulset, metav1.CreateOptions{})
	if err != nil {
		zap.L().Error("C-CreateStatefulset 创建Statefulset失败, ", zap.Error(err))
		return errors.New("创建Statefulset失败" + err.Error())
	}

	return nil
}

func CreateVolumeClaim(data *StatefulsetCreate) []corev1.PersistentVolumeClaim {
	vcNames := strings.Split(data.VolumeClaimName, ",")
	storages := strings.Split(data.Storage, ",")
	accessmodes := strings.Split(data.AccessMode, "/")
	//这里共用一个accessmode
	corev1AccessMode := []corev1.PersistentVolumeAccessMode{}
	for i, val := range accessmodes {
		corev1AccessMode[i] = corev1.PersistentVolumeAccessMode(val)
	}
	pvc := []corev1.PersistentVolumeClaim{}
	for idx, vcName := range vcNames {
		pvcItem := corev1.PersistentVolumeClaim{
			ObjectMeta: metav1.ObjectMeta{
				Name: vcName,
			},
		}

		pvcItem.Spec.AccessModes = corev1AccessMode
		if idx < len(storages) && data.Storage != "" {
			pvcItem.Spec.Resources.Requests = map[corev1.ResourceName]resource.Quantity{
				corev1.ResourceStorage: resource.MustParse(storages[idx]),
			}
		}
		pvc = append(pvc, pvcItem)
	}
	return pvc
}

func CreateContainer(data *StatefulsetCreate) []corev1.Container {
	var (
		imgs           = strings.Split(data.Image, ",")
		LimitCpus      = strings.Split(data.LimitCpu, ",")
		LimitMems      = strings.Split(data.LimitMemory, ",")
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
			Name:  data.Name + imgs[i] + "-" + strconv.Itoa(i),
			Image: imgs[i],
			VolumeMounts: []corev1.VolumeMount{
				{
					Name:      data.Name,
					MountPath: data.MountPath,
				},
			},
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
