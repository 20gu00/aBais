package service

import (
	"context"
	"errors"

	dataDispose "github.com/20gu00/aBais/common/data-dispose"

	"go.uber.org/zap"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

var StatefulSet statefulSet

type statefulSet struct{}

type StatusfulSetsResp struct {
	Items []appsv1.StatefulSet `json:"items"`
	Total int                  `json:"total"`
}

// 获取statefulset列表，支持过滤、排序、分页
func (s *statefulSet) GetStatefulSets(client *kubernetes.Clientset, filterName, namespace string, limit, page int) (statusfulSetsResp *StatusfulSetsResp, err error) {
	// 获取statefulSetList类型的statefulSet列表
	statefulSetList, err := client.AppsV1().StatefulSets(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		zap.L().Error("S-GetStatefulSets 获取StatefulSet列表失败", zap.Error(err))
		return nil, errors.New("获取StatefulSet列表失败, " + err.Error())
	}
	// 将statefulSetList中的StatefulSet列表(Items)，放进dataselector对象中，进行排序
	selectableData := &dataDispose.DataSelector{
		GenericDataList: s.toCells(statefulSetList.Items),
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

	// 将[]DataCell类型的statefulset列表转为v1.statefulset列表
	statefulSets := s.fromCells(data.GenericDataList)

	return &StatusfulSetsResp{
		Items: statefulSets,
		Total: total,
	}, nil
}

func (s *statefulSet) toCells(std []appsv1.StatefulSet) []dataDispose.DataCell {
	cells := make([]dataDispose.DataCell, len(std))
	for i := range std {
		cells[i] = dataDispose.StatefulSetCell(std[i])
	}
	return cells
}

func (s *statefulSet) fromCells(cells []dataDispose.DataCell) []appsv1.StatefulSet {
	statefulSets := make([]appsv1.StatefulSet, len(cells))
	for i := range cells {
		statefulSets[i] = appsv1.StatefulSet(cells[i].(dataDispose.StatefulSetCell))
	}

	return statefulSets
}
