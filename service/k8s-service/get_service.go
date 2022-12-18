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

var K8sService k8sService

type k8sService struct{}

type ServicesResp struct {
	Items []corev1.Service `json:"items"`
	Total int              `json:"total"`
}

//  获取service列表 过滤、排序、分页
func (s *k8sService) GetServices(client *kubernetes.Clientset, filterName, namespace string, limit, page int) (servicesResp *ServicesResp, err error) {
	//获取serviceList类型的service列表
	serviceList, err := client.CoreV1().Services(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		zap.L().Error("C-GetServices 获取Service列表失败", zap.Error(err))
		return nil, errors.New("获取Service列表失败, " + err.Error())
	}
	//将serviceList中的service列表(Items)，放进dataselector对象中，进行排序
	selectableData := &dataDispose.DataSelector{
		GenericDataList: s.toCells(serviceList.Items),
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

	//将[]DataCell类型的service列表转为v1.service列表
	services := s.fromCells(data.GenericDataList)

	return &ServicesResp{
		Items: services,
		Total: total,
	}, nil
}

func (s *k8sService) toCells(std []corev1.Service) []dataDispose.DataCell {
	cells := make([]dataDispose.DataCell, len(std))
	for i := range std {
		cells[i] = dataDispose.ServiceCell(std[i])
	}
	return cells
}

func (s *k8sService) fromCells(cells []dataDispose.DataCell) []corev1.Service {
	services := make([]corev1.Service, len(cells))
	for i := range cells {
		services[i] = corev1.Service(cells[i].(dataDispose.ServiceCell))
	}

	return services
}
