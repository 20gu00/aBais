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

var RoleBinding roleBinding

type roleBinding struct{}

type RoleBindingResp struct {
	Items []rbacv1.RoleBinding `json:"items"`
	Total int                  `json:"total"`
}

// 获取secret列表 过滤、排序、分页
func (s *roleBinding) GetRoleBindings(client *kubernetes.Clientset, filterName, namespace string, limit, page int) (roleBindingResp *RoleBindingResp, err error) {
	// 获取secretList类型的secret列表
	roleBindingList, err := client.RbacV1().RoleBindings(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		zap.L().Error("S-GetRoleBindings 获取Role Binding列表失败", zap.Error(err))
		return nil, errors.New("获取Role Binding列表失败, " + err.Error())
	}
	// 将secretList中的secret列表(Items)，放进dataselector对象中，进行排序
	selectableData := &dataDispose.DataSelector{
		GenericDataList: s.toCells(roleBindingList.Items),
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

	return &RoleBindingResp{
		Items: roles,
		Total: total,
	}, nil
}

func (s *roleBinding) toCells(std []rbacv1.RoleBinding) []dataDispose.DataCell {
	cells := make([]dataDispose.DataCell, len(std))
	for i := range std {
		cells[i] = dataDispose.RoleBindingCell(std[i])
	}
	return cells
}

func (s *roleBinding) fromCells(cells []dataDispose.DataCell) []rbacv1.RoleBinding {
	roleBindings := make([]rbacv1.RoleBinding, len(cells))
	for i := range cells {
		roleBindings[i] = rbacv1.RoleBinding(cells[i].(dataDispose.RoleBindingCell))
	}

	return roleBindings
}
