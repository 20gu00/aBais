package service

import (
	"context"
	"errors"
	"k8s.io/apimachinery/pkg/api/resource"
	"strings"

	"go.uber.org/zap"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type PvcCreate struct {
	Name         string            `json:"name"`
	Namespace    string            `json:"namespace"`
	Cluster      string            `json:"cluster"`
	Data         map[string]string `json:"data"`
	StorageClass string            `json:"storage_class"`
	VolumeMode   string            `json:"volume_mode"`
	Storage      string            `json:"storage"`
	AccessMode   string            `json:"access_mode"`
}

func (p *pvc) CreatePvc(client *kubernetes.Clientset, data *PvcCreate) (err error) {
	amsTemp := strings.Split(data.AccessMode, "/")
	ams := []corev1.PersistentVolumeAccessMode{}
	for idx, val := range amsTemp {
		ams[idx] = corev1.PersistentVolumeAccessMode(val)
	}
	pvc := &corev1.PersistentVolumeClaim{
		ObjectMeta: metav1.ObjectMeta{
			Name:      data.Name,
			Namespace: data.Namespace,
		},
		Spec: corev1.PersistentVolumeClaimSpec{
			AccessModes: ams,
			Resources: corev1.ResourceRequirements{
				Requests: map[corev1.ResourceName]resource.Quantity{
					corev1.ResourceStorage: resource.MustParse(data.Storage),
				},
			},
			StorageClassName: &data.StorageClass,
		},
	}

	_, err = client.CoreV1().PersistentVolumeClaims(data.Namespace).Create(context.TODO(), pvc, metav1.CreateOptions{})
	if err != nil {
		zap.L().Error("C-CreatePvc 创建pvc失败, ", zap.Error(err))
		return errors.New("创建pvc失败" + err.Error())
	}

	return nil
}
