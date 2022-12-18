package pvc

import (
	"context"
	"errors"

	dataDispose "github.com/20gu00/aBais/common/data-dispose"

	"go.uber.org/zap"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

var Pvc pvc

type pvc struct{}

type PvcsResp struct {
	Items []corev1.PersistentVolumeClaim `json:"items"`
	Total int                            `json:"total"`
}

// 获取pvc列表 过滤、排序、分页
func (p *pvc) GetPvcs(client *kubernetes.Clientset, filterName, namespace string, limit, page int) (pvcsResp *PvcsResp, err error) {
	// 获取pvcList类型的pvc列表
	pvcList, err := client.CoreV1().PersistentVolumeClaims(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		zap.L().Error("S-GetPvcs 获取Pvc列表失败", zap.Error(err))
		return nil, errors.New("获取Pvc列表失败, " + err.Error())
	}

	// 将pvcList中的pvc列表(Items)，放进dataselector对象中，进行排序
	selectableData := &dataDispose.DataSelector{
		GenericDataList: p.toCells(pvcList.Items),
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

	//将[]DataCell类型的pvc列表转为v1.pvc列表
	pvcs := p.fromCells(data.GenericDataList)

	return &PvcsResp{
		Items: pvcs,
		Total: total,
	}, nil
}

func (p *pvc) toCells(std []corev1.PersistentVolumeClaim) []dataDispose.DataCell {
	cells := make([]dataDispose.DataCell, len(std))
	for i := range std {
		cells[i] = dataDispose.PvcCell(std[i])
	}
	return cells
}

func (p *pvc) fromCells(cells []dataDispose.DataCell) []corev1.PersistentVolumeClaim {
	pvcs := make([]corev1.PersistentVolumeClaim, len(cells))
	for i := range cells {
		pvcs[i] = corev1.PersistentVolumeClaim(cells[i].(dataDispose.PvcCell))
	}

	return pvcs
}
