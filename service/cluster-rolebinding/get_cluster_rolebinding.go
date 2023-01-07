package clusterRoleBinding

import (
	"context"
	"errors"

	dataDispose "github.com/20gu00/aBais/common/data-dispose"

	"go.uber.org/zap"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

var ClusterRoleBinding clusterRoleBinding

type clusterRoleBinding struct{}

type ClusterRoleBindingResp struct {
	Items []rbacv1.ClusterRoleBinding `json:"items"`
	Total int                         `json:"total"`
}

// 获取secret列表 过滤、排序、分页
func (s *clusterRoleBinding) GetClusterRoleBindings(client *kubernetes.Clientset, filterName string, limit, page int) (clusterRoleBindingResp *ClusterRoleBindingResp, err error) {
	// 获取secretList类型的secret列表
	clusterRoleBindingList, err := client.RbacV1().ClusterRoleBindings().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		zap.L().Error("S-GetClusterRoleBindings 获取ClusterRoleBinding列表失败", zap.Error(err))
		return nil, errors.New("获取ClusterRoleBinding列表失败, " + err.Error())
	}
	// 将secretList中的secret列表(Items)，放进dataselector对象中，进行排序
	selectableData := &dataDispose.DataSelector{
		GenericDataList: s.toCells(clusterRoleBindingList.Items),
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

	//将[]DataCell类型的secret列表转为v1.secret列表
	roles := s.fromCells(data.GenericDataList)

	return &ClusterRoleBindingResp{
		Items: roles,
		Total: total,
	}, nil
}

func (s *clusterRoleBinding) toCells(std []rbacv1.ClusterRoleBinding) []dataDispose.DataCell {
	cells := make([]dataDispose.DataCell, len(std))
	for i := range std {
		cells[i] = dataDispose.ClusterRoleBindingCell(std[i])
	}
	return cells
}

func (s *clusterRoleBinding) fromCells(cells []dataDispose.DataCell) []rbacv1.ClusterRoleBinding {
	clusterRolelBindings := make([]rbacv1.ClusterRoleBinding, len(cells))
	for i := range cells {
		clusterRolelBindings[i] = rbacv1.ClusterRoleBinding(cells[i].(dataDispose.ClusterRoleBindingCell))
	}

	return clusterRolelBindings
}
