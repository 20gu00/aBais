package helmClient

import (
	"errors"
	"fmt"
	"log"
	"os"

	k8sClient "github.com/20gu00/aBais/common/k8s-clientset"
	service "github.com/20gu00/aBais/service/namespace"

	"go.uber.org/zap"
	"helm.sh/helm/v3/pkg/action"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

var HelmConfig helmConfig

type helmConfig struct {
	//helm客户端配置
	ActionConfigMap map[string]*action.Configuration
}

//no init
func (h *helmConfig) Init() {
	mp := make(map[string]*action.Configuration, 0)
	for cluster, kubeconfig := range k8sClient.K8s.KubeConfMap {
		client := k8sClient.K8s.ClientMap[cluster]
		namespaces, err := service.Namespace.GetNamespaces(client, "", 0, 0)
		if err != nil {
			panic(err)
		}
		for _, namespace := range namespaces.Items {
			actionConfig := new(action.Configuration)
			cf := genericclioptions.ConfigFlags{
				KubeConfig: &kubeconfig,
				Namespace:  &namespace.Name,
			}
			//namespace为空字符串那么就是列出全部namespace
			if err := actionConfig.Init(&cf, namespace.Name, os.Getenv("HELM_DRIVER"), log.Printf); err != nil {
				zap.L().Error("helm Config初始化失败", zap.Error(err))
				panic("helm Config初始化失败, " + err.Error())
			}
			str := fmt.Sprintf("%s-%s", namespace.Name, cluster)
			mp[str] = actionConfig
			zap.L().Info(fmt.Sprintf("集群:%s,命名空间:%s,初始化action Config成功 ", cluster, namespace.Name))
		}
	}
	h.ActionConfigMap = mp
}

func (*helmConfig) GetAc(cluster, namespace string) (*action.Configuration, error) {
	kubeconfig, ok := K8s.KubeConfMap[cluster]
	if !ok {
		logger.Error("actionConfig初始化失败, cluster不存在")
		return nil, errors.New("actionConfig初始化失败, cluster不存在")
	}
	actionConfig := new(action.Configuration)
	cf := &genericclioptions.ConfigFlags{
		KubeConfig: &kubeconfig,
		Namespace:  &namespace,
	}
	if err := actionConfig.Init(cf, namespace, os.Getenv("HELM_DRIVER"), log.Printf); err != nil {
		logger.Error("actionConfig初始化失败, %+v", err)
		return nil, errors.New("actionConfig初始化失败, " + err.Error())
	}
	return actionConfig, nil
}
