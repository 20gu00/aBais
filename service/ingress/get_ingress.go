package service

import (
	"context"
	"errors"

	dataDispose "github.com/20gu00/aBais/common/data-dispose"

	"go.uber.org/zap"
	nwv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

var Ingress ingress

type ingress struct{}

type IngressesResp struct {
	Items []nwv1.Ingress `json:"items"`
	Total int            `json:"total"`
}

// 获取ingress列表过滤、排序、分页
func (i *ingress) GetIngresses(client *kubernetes.Clientset, filterName, namespace string, limit, page int) (ingressesResp *IngressesResp, err error) {
	// 获取ingressList类型的ingress列表
	ingressList, err := client.NetworkingV1().Ingresses(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		zap.L().Error("S-GetIngresses 获取Ingress列表失败", zap.Error(err))
		return nil, errors.New("获取Ingress列表失败, " + err.Error())
	}
	// 将ingressList中的ingress列表(Items)，放进dataselector对象中，进行排序
	selectableData := &dataDispose.DataSelector{
		GenericDataList: i.toCells(ingressList.Items),
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

	//将[]DataCell类型的ingress列表转为v1.ingress列表
	ingresss := i.fromCells(data.GenericDataList)

	return &IngressesResp{
		Items: ingresss,
		Total: total,
	}, nil
}

func (i *ingress) toCells(std []nwv1.Ingress) []dataDispose.DataCell {
	cells := make([]dataDispose.DataCell, len(std))
	for i := range std {
		cells[i] = dataDispose.IngressCell(std[i])
	}
	return cells
}

func (i *ingress) fromCells(cells []dataDispose.DataCell) []nwv1.Ingress {
	ingresss := make([]nwv1.Ingress, len(cells))
	for i := range cells {
		ingresss[i] = nwv1.Ingress(cells[i].(dataDispose.IngressCell))
	}

	return ingresss
}
