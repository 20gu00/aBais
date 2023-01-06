package service

import (
	"context"
	"errors"
	batchv1 "k8s.io/api/batch/v1"
	"strings"

	"go.uber.org/zap"
	batchv1beta1 "k8s.io/api/batch/v1beta1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type CronJobCreate struct {
	Name          string `json:"name"`
	Namespace     string `json:"namespace"`
	Cluster       string `json:"cluster"`
	RestartPolicy string `json:"restart_policy"`
	Image         string `json:"image"`
	Schedule      string `json:"schedule"`
	Command       string `json:"command"`
}

func (c *cronjob) CreateCronJob(client *kubernetes.Clientset, data *CronJobCreate) (err error) {
	command := strings.Split(data.Command, "|")
	cronJob := &batchv1beta1.CronJob{
		ObjectMeta: metav1.ObjectMeta{
			Name:      data.Name,
			Namespace: data.Namespace,
		},
		Spec: batchv1beta1.CronJobSpec{
			Schedule: data.Schedule,
			JobTemplate: batchv1beta1.JobTemplateSpec{
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
			},
		},
	}

	_, err = client.BatchV1beta1().CronJobs(data.Namespace).Create(context.TODO(), cronJob, metav1.CreateOptions{})
	if err != nil {
		zap.L().Error("S-CreateCronJob 创建CronJob失败, ", zap.Error(err))
		return errors.New("创建CronJob失败" + err.Error())
	}

	return nil
}
