package service

import (
	"context"
	"errors"

	dataDispose "github.com/20gu00/aBais/common/data-dispose"

	"go.uber.org/zap"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

var Role role

type role struct{}

type RoleResp struct {
	Items []rbacv1.Role `json:"items"`
	Total int           `json:"total"`
}

// 获取secret列表 过滤、排序、分页
func (s *role) GetRoles(client *kubernetes.Clientset, filterName, namespace string, limit, page int) (roleResp *RoleResp, err error) {
	// 获取secretList类型的secret列表
	roleList, err := client.RbacV1().Roles(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		zap.L().Error("S-GetRoles 获取Role列表失败", zap.Error(err))
		return nil, errors.New("获取Role列表失败, " + err.Error())
	}
	// 将secretList中的secret列表(Items)，放进dataselector对象中，进行排序
	selectableData := &dataDispose.DataSelector{
		GenericDataList: s.toCells(roleList.Items),
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

	return &RoleResp{
		Items: roles,
		Total: total,
	}, nil
}

func (s *role) toCells(std []rbacv1.Role) []dataDispose.DataCell {
	cells := make([]dataDispose.DataCell, len(std))
	for i := range std {
		cells[i] = dataDispose.RoleCell(std[i])
	}
	return cells
}

func (s *role) fromCells(cells []dataDispose.DataCell) []rbacv1.Role {
	roles := make([]rbacv1.Role, len(cells))
	for i := range cells {
		roles[i] = rbacv1.Role(cells[i].(dataDispose.RoleCell))
	}

	return roles
}
