package service

import (
	"context"
	"errors"
	"strings"

	"go.uber.org/zap"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type JobCreate struct {
	Name          string `json:"name"`
	Namespace     string `json:"namespace"`
	Cluster       string `json:"cluster"`
	RestartPolicy string `json:"restart_policy"`
	Image         string `json:"image"`
	Command       string `json:"command"`
}

func (c *job) CreateJob(client *kubernetes.Clientset, data *JobCreate) (err error) {
	command := strings.Split(data.Command, "|")
	job := &batchv1.Job{
		ObjectMeta: metav1.ObjectMeta{
			Name:      data.Name,
			Namespace: data.Namespace,
		},
		Spec: batchv1.JobSpec{
			Template: corev1.PodTemplateSpec{
				Spec: corev1.PodSpec{
					RestartPolicy: corev1.RestartPolicy(data.RestartPolicy),
					Containers: []corev1.Container{
						{
							Name:    data.Name + data.Image,
							Image:   data.Image,
							Command: command,
						},
					},
				},
			},
		},
	}

	_, err = client.BatchV1().Jobs(data.Namespace).Create(context.TODO(), job, metav1.CreateOptions{})
	if err != nil {
		zap.L().Error("S-CreateJob 创建Job失败, ", zap.Error(err))
		return errors.New("创建Job失败" + err.Error())
	}

	return nil
}