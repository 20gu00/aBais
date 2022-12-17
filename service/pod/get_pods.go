package pod

import (
	"context"
	"errors"

	dataDispose "github.com/20gu00/aBais/common/data-dispose"

	"go.uber.org/zap"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

var Pod pod

type pod struct{}

type PodsResp struct {
	Items []corev1.Pod `json:"items"`
	Total int          `json:"total"`
}

func (p *pod) GetPods(client *kubernetes.Clientset, filterName, namespace string, limit, page int) (podsResp *PodsResp, err error) {
	// 获取podList类型的pod列表
	podList, err := client.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		zap.L().Error("S-GetPods 获取Pod列表失败, ", zap.Error(err))
		return nil, errors.New("获取Pod列表失败, " + err.Error())
	}
	// 实例化dataSelector对象
	selectableData := &dataDispose.DataSelector{
		GenericDataList: p.toCells(podList.Items),
		DataSelectQuery: &dataDispose.DataSelectQuery{
			FilterQuery: &dataDispose.FilterQuery{Name: filterName},
			PaginateQuery: &dataDispose.PaginateQuery{
				Limit: limit,
				Page:  page,
			},
		},
	}

	// 1.过滤
	filtered := selectableData.Filter()
	total := len(filtered.GenericDataList) // pod数目
	//2.排序 分页
	data := filtered.Sort().Paginate()

	//将[]DataCell类型的pod列表转为v1.pod列表
	pods := p.fromCells(data.GenericDataList)

	return &PodsResp{
		Items: pods,
		Total: total,
	}, nil
}

// 将pod类型切片，转换成DataCell类型切片
func (p *pod) toCells(std []corev1.Pod) []dataDispose.DataCell {
	cells := make([]dataDispose.DataCell, len(std))
	for i := range std {
		cells[i] = dataDispose.PodCell(std[i]) // 数据类型强转前提
	}
	return cells
}

// []DataCell -> []pod
func (p *pod) fromCells(cells []dataDispose.DataCell) []corev1.Pod {
	pods := make([]corev1.Pod, len(cells))
	for i := range cells {
		pods[i] = corev1.Pod(cells[i].(dataDispose.PodCell))
	}
	return pods
}
