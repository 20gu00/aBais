package all_resource

import (
	"context"
	"sync"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

var AllRes allRes

type allRes struct{}

// 获取集群所有资源数量
func (a *allRes) GetAllResourceNum(client *kubernetes.Clientset) (map[string]int, []error) {
	var wg sync.WaitGroup
	wg.Add(12)

	errs := make([]error, 0)
	data := make(map[string]int, 0)
	go func() {
		list, err := client.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			errs = append(errs, err)
		}
		data["Nodes"] = len(list.Items)
		wg.Done()
	}()
	go func() {
		list, err := client.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			errs = append(errs, err)
		}
		data["Namespaces"] = len(list.Items)
		wg.Done()
	}()
	go func() {
		list, err := client.NetworkingV1().Ingresses("").List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			errs = append(errs, err)
		}
		data["Ingresses"] = len(list.Items)
		wg.Done()
	}()
	go func() {
		list, err := client.CoreV1().PersistentVolumes().List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			errs = append(errs, err)
		}
		data["PVs"] = len(list.Items)
		wg.Done()
	}()
	go func() {
		list, err := client.AppsV1().DaemonSets("").List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			errs = append(errs, err)
		}
		data["DaemonSets"] = len(list.Items)
		wg.Done()
	}()
	go func() {
		list, err := client.AppsV1().StatefulSets("").List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			errs = append(errs, err)
		}
		data["StatefulSets"] = len(list.Items)
		wg.Done()
	}()
	go func() {
		list, err := client.BatchV1().Jobs("").List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			errs = append(errs, err)
		}
		data["Jobs"] = len(list.Items)
		wg.Done()
	}()
	go func() {
		list, err := client.CoreV1().Services("").List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			errs = append(errs, err)
		}
		data["Services"] = len(list.Items)
		wg.Done()
	}()
	go func() {
		list, err := client.AppsV1().Deployments("").List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			errs = append(errs, err)
		}
		data["Deployments"] = len(list.Items)
		wg.Done()
	}()
	go func() {
		list, err := client.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			errs = append(errs, err)
		}
		data["Pods"] = len(list.Items)
		wg.Done()
	}()
	go func() {
		list, err := client.CoreV1().Secrets("").List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			errs = append(errs, err)
		}
		data["Secrets"] = len(list.Items)
		wg.Done()
	}()
	go func() {
		list, err := client.CoreV1().ConfigMaps("").List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			errs = append(errs, err)
		}
		data["ConfigMaps"] = len(list.Items)
		wg.Done()
	}()
	wg.Wait()
	return data, nil
}
