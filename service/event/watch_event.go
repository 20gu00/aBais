package service

import (
	"fmt"
	"time"

	"github.com/20gu00/aBais/dao"
	"github.com/20gu00/aBais/model"

	k8sClient "github.com/20gu00/aBais/common/k8s-clientset"
	"go.uber.org/zap"
	coreV1 "k8s.io/api/core/v1"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/tools/cache"
)

//监听event
func (*event) WatchEventTask(cluster string) {
	// 不通过client-go的watch而是informer机制去watchk8s集群资源,工厂模式获取informer(event频繁)
	informerFactory := informers.NewSharedInformerFactory(k8sClient.K8s.ClientMap[cluster], time.Minute)
	// 要处理(reflector的list watch)的资源对象类型
	informer := informerFactory.Core().V1().Events()
	// Informer()类似于注册进工厂,要在start前
	// 添加事件的回调函数
	informer.Informer().AddEventHandler(
		cache.ResourceEventHandlerFuncs{
			AddFunc: func(obj interface{}) {
				onAdd(obj, cluster)
			},
			//UpdateFunc: onUpdate,
			//DeleteFunc: onDelete,
		},
	)
	// 关闭时工厂关闭
	stopCh := make(chan struct{})
	defer close(stopCh)
	// 工厂开启(或者goroutine,多个协程可以是使用同一个informer,即不用太多reflector去和apiserver交互)
	informerFactory.Start(stopCh)
	// 判断本地缓存和etcd数据是否同步
	if !cache.WaitForCacheSync(stopCh, informer.Informer().HasSynced) {
		zap.L().Error("等房贷本地缓存同步超时")
		return
	}
	<-stopCh

	return
}

// 将事件存入数据库
func onAdd(obj interface{}, cluster string) {
	// 断言event(involve涉及)
	event := obj.(*coreV1.Event)
	_, has, err := dao.Event.HasEvent(event.InvolvedObject.Name,
		event.InvolvedObject.Kind,
		event.InvolvedObject.Namespace,
		event.Reason,
		event.CreationTimestamp.Time,
		cluster,
	)

	if err != nil {
		return
	}
	if has {
		zap.L().Error(fmt.Sprintf("Event数据已存在, %s %s %s %s %v %s\n",
			event.InvolvedObject.Name,
			event.InvolvedObject.Kind,
			event.InvolvedObject.Namespace,
			event.Reason,
			event.CreationTimestamp.Time,
			cluster),
		)
		return
	}
	data := &model.K8sEvent{
		Name:      event.InvolvedObject.Name,
		Kind:      event.InvolvedObject.Kind,
		Namespace: event.InvolvedObject.Namespace,
		Rtype:     event.Type,
		Reason:    event.Reason,
		Message:   event.Message,
		EventTime: &event.CreationTimestamp.Time,
		Cluster:   cluster,
	}
	if err := dao.Event.Add(data); err != nil {
		return
	}
}
