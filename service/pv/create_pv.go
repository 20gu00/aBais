package service

import (
	"context"
	"errors"
	"strings"

	"go.uber.org/zap"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type PvCreate struct {
	Name         string `json:"name"`
	Cluster      string `json:"cluster"`
	StorageClass string `json:"storage_class"`
	VolumeMode   string `json:"volume_mode"`
	Storage      string `json:"storage"`
	AccessMode   string `json:"access_mode"`
	Path         string `json:"path"`
}

func (p *pv) CreatePvc(client *kubernetes.Clientset, data *PvCreate) (err error) {
	amsTemp := strings.Split(data.AccessMode, "/")
	ams := []corev1.PersistentVolumeAccessMode{}
	//ams := make([]corev1.PersistentVolumeAccessMode, 10)  这两者长度都默认是0

	//var vmTemp corev1.PersistentVolumeMode
	vmTemp := corev1.PersistentVolumeMode(data.VolumeMode)
	for _, val := range amsTemp {
		//ams[idx] = corev1.PersistentVolumeAccessMode(val)  //得有长度
		ams = append(ams, corev1.PersistentVolumeAccessMode(val))
	}
	pv := &corev1.PersistentVolume{
		ObjectMeta: metav1.ObjectMeta{
			Name: data.Name,
		},
		Spec: corev1.PersistentVolumeSpec{
			AccessModes: ams,
			VolumeMode:  &vmTemp,
			PersistentVolumeSource: corev1.PersistentVolumeSource{
				HostPath: &corev1.HostPathVolumeSource{
					Path: data.Path,
				},
			},
			Capacity: corev1.ResourceList{
				corev1.ResourceStorage: resource.MustParse(data.Storage),
			},
			//Resources: corev1.ResourceRequirements{
			//	Requests: map[corev1.ResourceName]resource.Quantity{
			//		corev1.ResourceStorage: resource.MustParse(data.Storage),
			//	},
			//},
			StorageClassName: data.StorageClass,
		},
	}

	_, err = client.CoreV1().PersistentVolumes().Create(context.TODO(), pv, metav1.CreateOptions{})
	if err != nil {
		zap.L().Error("C-CreatePv 创建pv失败, ", zap.Error(err))
		return errors.New("创建pv失败" + err.Error())
	}

	return nil
}
