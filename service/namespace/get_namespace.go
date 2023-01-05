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

var Namespace namespace

type namespace struct{}

type NamespacesResp struct {
	Items []corev1.Namespace `json:"items"`
	Total int                `json:"total"`
}

// 获取namespace列表，支持过滤、排序、分页
func (n *namespace) GetNamespaces(client *kubernetes.Clientset, filterName string, limit, page int) (namespacesResp *NamespacesResp, err error) {
	// 获取namespaceList类型的namespace列表
	namespaceList, err := client.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		zap.L().Error("S-GetNamespaces 获取Namespace列表失败", zap.Error(err))
		return nil, errors.New("获取Namespace列表失败" + err.Error())
	}

	// 将namespaceList中的namespace列表(Items)，放进dataselector对象中，进行排序
	selectableData := &dataDispose.DataSelector{
		GenericDataList: n.toCells(namespaceList.Items),
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

	//将[]DataCell类型的namespace列表转为v1.namespace列表
	namespaces := n.fromCells(data.GenericDataList)

	return &NamespacesResp{
		Items: namespaces,
		Total: total,
	}, nil
}

func (n *namespace) toCells(std []corev1.Namespace) []dataDispose.DataCell {
	cells := make([]dataDispose.DataCell, len(std))
	for i := range std {
		cells[i] = dataDispose.NamespaceCell(std[i])
	}
	return cells
}

func (n *namespace) fromCells(cells []dataDispose.DataCell) []corev1.Namespace {
	namespaces := make([]corev1.Namespace, len(cells))
	for i := range cells {
		namespaces[i] = corev1.Namespace(cells[i].(dataDispose.NamespaceCell))
	}

	return namespaces
}
