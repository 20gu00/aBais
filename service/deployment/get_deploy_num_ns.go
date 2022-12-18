package service

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// 定义DeploysNs类型，用于返回namespace中deployment的数量
type DeploysNs struct {
	Namespace string `json:"namespace"`
	DeployNum int    `json:"deployment_num"`
}

// 获取每个namespace的deployment数量
func (d *deployment) GetDeployNumPerNs(client *kubernetes.Clientset) (deploysNps []*DeploysNs, err error) {
	namespaceList, err := client.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	for _, namespace := range namespaceList.Items {
		deploymentList, err := client.AppsV1().Deployments(namespace.Name).List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			return nil, err
		}

		deploysNp := &DeploysNs{
			Namespace: namespace.Name,
			DeployNum: len(deploymentList.Items),
		}

		deploysNps = append(deploysNps, deploysNp)
	}
	return deploysNps, nil
}
