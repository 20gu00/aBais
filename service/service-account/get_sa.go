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

var Sa sa

type sa struct{}

type SasResp struct {
	Items []corev1.ServiceAccount `json:"items"`
	Total int                     `json:"total"`
}

func (j *sa) GetSas(client *kubernetes.Clientset, filterName, namespace string, limit, page int) (sasResp *SasResp, err error) {
	saList, err := client.CoreV1().ServiceAccounts(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		zap.L().Error("S-GetSas 获取sa列表失败", zap.Error(err))
		return nil, errors.New("获取sa列表失败, " + err.Error())
	}

	selectableData := &dataDispose.DataSelector{
		GenericDataList: j.toCells(saList.Items),
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

	jobs := j.fromCells(data.GenericDataList)

	return &SasResp{
		Items: jobs,
		Total: total,
	}, nil
}

func (j *sa) toCells(std []corev1.ServiceAccount) []dataDispose.DataCell {
	cells := make([]dataDispose.DataCell, len(std))
	for i := range std {
		cells[i] = dataDispose.ServiceAccountCell(std[i])
	}
	return cells
}

func (j *sa) fromCells(cells []dataDispose.DataCell) []corev1.ServiceAccount {
	sas := make([]corev1.ServiceAccount, len(cells))
	for i := range cells {
		sas[i] = corev1.ServiceAccount(cells[i].(dataDispose.ServiceAccountCell))
	}

	return sas
}
