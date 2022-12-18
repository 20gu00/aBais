package service

import (
	"context"
	"errors"

	dataDispose "github.com/20gu00/aBais/common/data-dispose"

	"go.uber.org/zap"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

var Pv pv

type pv struct{}

type PvsResp struct {
	Items []corev1.PersistentVolume `json:"items"`
	Total int                       `json:"total"`
}

// 获取pv列表 过滤、排序、分页
func (p *pv) GetPvs(client *kubernetes.Clientset, filterName string, limit, page int) (pvsResp *PvsResp, err error) {
	// 获取pvList类型的pv列表
	pvList, err := client.CoreV1().PersistentVolumes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		zap.L().Error("S-GetPvs 获取Pv列表失败", zap.Error(err))
		return nil, errors.New("获取Pv列表失败, " + err.Error())
	}

	// 将pvList中的pv列表(Items)，放进dataselector对象中，进行排序
	selectableData := &dataDispose.DataSelector{
		GenericDataList: p.toCells(pvList.Items),
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

	// 将[]DataCell类型的pv列表转为v1.pv列表
	pvs := p.fromCells(data.GenericDataList)

	return &PvsResp{
		Items: pvs,
		Total: total,
	}, nil
}

func (p *pv) toCells(std []corev1.PersistentVolume) []dataDispose.DataCell {
	cells := make([]dataDispose.DataCell, len(std))
	for i := range std {
		cells[i] = dataDispose.PvCell(std[i])
	}
	return cells
}

func (p *pv) fromCells(cells []dataDispose.DataCell) []corev1.PersistentVolume {
	pvs := make([]corev1.PersistentVolume, len(cells))
	for i := range cells {
		pvs[i] = corev1.PersistentVolume(cells[i].(dataDispose.PvCell))
	}

	return pvs
}
