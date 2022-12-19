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

var Node node

type node struct{}

type NodesResp struct {
	Items []corev1.Node `json:"items"`
	Total int           `json:"total"`
}

// 获取node列表，支持过滤、排序、分页
func (n *node) GetNodes(client *kubernetes.Clientset, filterName string, limit, page int) (nodesResp *NodesResp, err error) {
	// 获取nodeList类型的node列表
	nodeList, err := client.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		zap.L().Error("S-GetNodes 获取Node列表失败", zap.Error(err))
		return nil, errors.New("获取Node列表失败, " + err.Error())
	}

	// 将nodeList中的node列表(Items)，放进dataselector对象中，进行排序
	selectableData := &dataDispose.DataSelector{
		GenericDataList: n.toCells(nodeList.Items),
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

	//将[]DataCell类型的node列表转为v1.node列表
	nodes := n.fromCells(data.GenericDataList)

	return &NodesResp{
		Items: nodes,
		Total: total,
	}, nil
}

func (n *node) toCells(std []corev1.Node) []dataDispose.DataCell {
	cells := make([]dataDispose.DataCell, len(std))
	for i := range std {
		cells[i] = dataDispose.NodeCell(std[i])
	}
	return cells
}

func (n *node) fromCells(cells []dataDispose.DataCell) []corev1.Node {
	nodes := make([]corev1.Node, len(cells))
	for i := range cells {
		nodes[i] = corev1.Node(cells[i].(dataDispose.NodeCell))
	}

	return nodes
}
