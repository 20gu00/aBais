package daemonset

import (
	"context"
	"errors"

	dataDispose "github.com/20gu00/aBais/common/data_dispose"

	"go.uber.org/zap"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

var DaemonSet daemonSet

type daemonSet struct{}

type DaemonSetsResp struct {
	Items []appsv1.DaemonSet `json:"items"`
	Total int                `json:"total"`
}

// 获取daemonset列表 过滤、排序、分页
func (d *daemonSet) GetDaemonSets(client *kubernetes.Clientset, filterName, namespace string, limit, page int) (daemonSetsResp *DaemonSetsResp, err error) {
	// 获取daemonSetList类型的daemonSet列表
	daemonSetList, err := client.AppsV1().DaemonSets(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		zap.L().Error("C-GetDaemonSets 获取DaemonSet列表失败, ", zap.Error(err))
		return nil, errors.New("获取DaemonSet列表失败, " + err.Error())
	}
	// 将daemonSetList中的daemonSet列表(Items)，放进dataselector对象中，进行排序
	selectableData := &dataDispose.DataSelector{
		GenericDataList: d.toCells(daemonSetList.Items),
		DataSelectQuery: &dataDispose.DataSelectQuery{
			FilterQuery: &dataDispose.FilterQuery{Name: filterName},
			PaginateQuery: &dataDispose.PaginateQuery{
				Limit: limit,
				Page:  page,
			},
		},
	}

	filtered := selectableData.Filter()
	total := len(filtered.GenericDataList)
	data := filtered.Sort().Paginate()

	// 将[]DataCell类型的daemonset列表转为v1.daemonset列表
	daemonSets := d.fromCells(data.GenericDataList)

	return &DaemonSetsResp{
		Items: daemonSets,
		Total: total,
	}, nil
}

func (d *daemonSet) toCells(std []appsv1.DaemonSet) []dataDispose.DataCell {
	cells := make([]dataDispose.DataCell, len(std))
	for i := range std {
		cells[i] = dataDispose.DaemonSetCell(std[i])
	}
	return cells
}

func (d *daemonSet) fromCells(cells []dataDispose.DataCell) []appsv1.DaemonSet {
	daemonSets := make([]appsv1.DaemonSet, len(cells))
	for i := range cells {
		daemonSets[i] = appsv1.DaemonSet(cells[i].(dataDispose.DaemonSetCell))
	}

	return daemonSets
}
