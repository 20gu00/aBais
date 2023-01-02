package service

import (
	"context"
	"errors"
	"go.uber.org/zap"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/kubernetes"
	"strconv"
	"strings"
)

// 定义DeployCreate结构体
type DeployCreate struct {
	Name          string            `json:"name"`
	Namespace     string            `json:"namespace"`
	Replicas      int32             `json:"replicas"`
	Image         string            `json:"image"`
	Label         map[string]string `json:"label"`
	PodLabel      map[string]string `json:"pod_label"`
	LimitCpu      string            `json:"limit_cpu"`
	LimitMemory   string            `json:"limit_memory"`
	RequestCpu    string            `json:"request_cpu"`
	RequestMemory string            `json:"request_memory"`
	//Cpu           string            `json:"cpu"`
	//Memory        string            `json:"memory"`
	ContainerPort string `json:"container_port"`
	HealthCheck   bool   `json:"health_check"`
	HealthPath    string `json:"health_path"`
	Cluster       string `json:"cluster"`
}

// 创建deployment
func (d *deployment) CreateDeployment(client *kubernetes.Clientset, data *DeployCreate) (err error) {
	deployment := &appsv1.Deployment{
		// GVK
		// ObjectMeta  metadata
		ObjectMeta: metav1.ObjectMeta{
			Name:      data.Name,
			Namespace: data.Namespace,
			Labels:    data.Label,
		},

		// Spec   deployment spec
		Spec: appsv1.DeploymentSpec{
			Replicas: &data.Replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: data.PodLabel,
			},
			// pod template
			Template: corev1.PodTemplateSpec{
				// 定义pod metadata
				ObjectMeta: metav1.ObjectMeta{
					Name:   data.Name,
					Labels: data.PodLabel,
				},
				// pod spec
				Spec: corev1.PodSpec{
					Containers: CreateContainer(data),
				},
			},
		},
		// Status定义资源的运行状态，这里由于是新建，传入空的appsv1.DeploymentStatus{}对象即可
		// 之后由kubernetes接管
		Status: appsv1.DeploymentStatus{},
	}

	// 探针probe 判断是否打开健康检查功能，若打开，则定义ReadinessProbe和LivenessProbe(startProbe)
	//if data.HealthCheck {
	//	// 设置第一个容器的ReadinessProbe，因为这里pod中只有一个容器，所以直接使用index 0即可 若pod中有多个容器，则这里需要使用for循环去定义了
	//	deployment.Spec.Template.Spec.Containers[0].ReadinessProbe = &corev1.Probe{
	//		ProbeHandler: corev1.ProbeHandler{
	//			// exec httpget tcp grpc
	//			// ping
	//			HTTPGet: &corev1.HTTPGetAction{
	//				Path: data.HealthPath,
	//				// port int string
	//				// intstr.IntOrString的作用是端口可以定义为整型，也可以定义为字符串
	//				// Type=0则表示表示该结构体实例内的数据为整型，转json时只使用IntVal的数据
	//				// Type=1则表示表示该结构体实例内的数据为字符串，转json时只使用StrVal的数据
	//				Port: intstr.IntOrString{
	//					Type:   0,
	//					IntVal: data.ContainerPort,
	//				},
	//			},
	//		},
	//		//初始化等待时间
	//		InitialDelaySeconds: 5,
	//		//超时时间
	//		TimeoutSeconds: 5,
	//		//执行间隔
	//		PeriodSeconds: 5,
	//		//// 探测失败后连续探测成功几次才算成功
	//		//SuccessThreshold: 3,
	//		//// 连续探测失败几次才算失败
	//		//FailureThreshold:1,
	//	}
	//}
	// 调用sdk(client cli)创建deployment
	_, err = client.AppsV1().Deployments(data.Namespace).Create(context.TODO(), deployment, metav1.CreateOptions{})
	if err != nil {
		zap.L().Error("C-CreateDeployment 创建Deployment失败, ", zap.Error(err))
		return errors.New("创建Deployment失败" + err.Error())
	}

	return nil
}

func CreateContainer(data *DeployCreate) []corev1.Container {
	// image limit request对应
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
		//不能直接用string转换整形
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
