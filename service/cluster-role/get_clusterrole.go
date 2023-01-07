package cluster_role

import (
	"context"
	"errors"

	dataDispose "github.com/20gu00/aBais/common/data-dispose"

	"go.uber.org/zap"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

var ClusterRole clusterRole

type clusterRole struct{}

type ClusterRoleResp struct {
	Items []rbacv1.ClusterRole `json:"items"`
	Total int                  `json:"total"`
}

// 获取secret列表 过滤、排序、分页
func (s *clusterRole) GetClusterRoles(client *kubernetes.Clientset, filterName string, limit, page int) (clusterRoleResp *ClusterRoleResp, err error) {
	// 获取secretList类型的secret列表
	clusterRoleList, err := client.RbacV1().ClusterRoles().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		zap.L().Error("S-GetClusterRoles 获取ClusterRole列表失败", zap.Error(err))
		return nil, errors.New("获取ClusterRole列表失败, " + err.Error())
	}
	// 将secretList中的secret列表(Items)，放进dataselector对象中，进行排序
	selectableData := &dataDispose.DataSelector{
		GenericDataList: s.toCells(clusterRoleList.Items),
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

	return &ClusterRoleResp{
		Items: roles,
		Total: total,
	}, nil
}

func (s *clusterRole) toCells(std []rbacv1.ClusterRole) []dataDispose.DataCell {
	cells := make([]dataDispose.DataCell, len(std))
	for i := range std {
		cells[i] = dataDispose.ClusterRoleCell(std[i])
	}
	return cells
}

func (s *clusterRole) fromCells(cells []dataDispose.DataCell) []rbacv1.ClusterRole {
	clusterRoles := make([]rbacv1.ClusterRole, len(cells))
	for i := range cells {
		clusterRoles[i] = rbacv1.ClusterRole(cells[i].(dataDispose.ClusterRoleCell))
	}

	return clusterRoles
}
