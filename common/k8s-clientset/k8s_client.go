package k8sClient

import (
	"encoding/json"
	"errors"
	"fmt"
	"sync"

	"github.com/20gu00/aBais/common/config"

	"go.uber.org/zap"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// 直接通过var实例化 k8s{xxx}  (Newxxx)
var K8s k8s

type k8s struct {
	// 多个客户端
	ClientMap map[string]*kubernetes.Clientset
	// 集群配置文件
	KubeConfMap map[string]string // 或者使用slice map查询快
	// 原生map加锁
	Locker sync.Mutex // sync.RWMutex
}

func (k *k8s) InitK8s() {
	k.Locker.Lock()
	defer k.Locker.Unlock()

	k.ClientMap = map[string]*kubernetes.Clientset{}

	// 将集群的配置文件路径信息反序列化到kubeMap string->map(json反序列化,将信息解码写入到比如结构体中)
	// map[Cluster-1:/root/.kube/config Cluster-2:/root/.kube/config]
	if err := json.Unmarshal([]byte(config.Config.KubeConfigs), &k.KubeConfMap); err != nil {
		panic(fmt.Sprintf("获取kubeConfig配置文件路径信息,Kubeconfigs反序列化失败 %v\n", err))
	}

	for key, v := range k.KubeConfMap {
		// 根据路径生成集群的配置文件(集群外或者集群内 Pod)
		conf, err := rest.InClusterConfig()
		if err != nil {
			conf, err = clientcmd.BuildConfigFromFlags("", v)
			if err != nil {
				panic(fmt.Sprintf("集群%s: 创建K8s配置失败 %v\n", key, err))
			}
		}

		// 根据配置文件生成clientset
		clientSet, err := kubernetes.NewForConfig(conf)
		if err != nil {
			panic(fmt.Sprintf("集群%s: 创建K8s Client失败 %v\n", key, err))
		}

		// key是cluster_name
		k.ClientMap[key] = clientSet
		zap.L().Info(fmt.Sprintf("集群%s: 创建K8s Client成功 ", key))
	}

}

// 获取操作集群的client
func (k *k8s) GetK8sClient(clusterName string) (*kubernetes.Clientset, error) {
	k.Locker.Lock()
	defer k.Locker.Unlock()

	k8sClient, ok := k.ClientMap[clusterName]
	if !ok {
		zap.L().Error(fmt.Sprintf("%s集群不存在", clusterName))
		return nil, errors.New(fmt.Sprintf("集群:%s 不存在, 无法获取client", clusterName))
	}

	return k8sClient, nil
}
