package deployment

import (
	"context"
	"errors"

	dataDispose "github.com/20gu00/aBais/common/data_dispose"

	"go.uber.org/zap"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

var Deployment deployment

type deployment struct{}

// Items是deployment元素列表，Total为deployment元素数量
type DeploymentsResp struct {
	Items []appsv1.Deployment `json:"items"`
	Total int                 `json:"total"`
}

// 获取deployment列表，支持过滤、排序、分页
func (d *deployment) GetDeployments(client *kubernetes.Clientset, filterName, namespace string, limit, page int) (deploymentsResp *DeploymentsResp, err error) {
	// 获取deploymentList类型的deployment列表(list)
	deploymentList, err := client.AppsV1().Deployments(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		zap.L().Error("S-GetDeployments 获取Deployment列表失败", zap.Error(err))
		return nil, errors.New("获取Deployment列表失败, " + err.Error())
	}

	// 将deploymentList中的deployment列表(Items)，放进dataselector对象中，进行排序
	selectableData := &dataDispose.DataSelector{
		GenericDataList: d.toCells(deploymentList.Items),
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

	//将[]DataCell类型的deployment列表转为appsv1.deployment列表
	deployments := d.fromCells(data.GenericDataList)

	return &DeploymentsResp{
		Items: deployments,
		Total: total,
	}, nil
}

func (d *deployment) toCells(std []appsv1.Deployment) []dataDispose.DataCell {
	cells := make([]dataDispose.DataCell, len(std))
	for i := range std {
		cells[i] = dataDispose.DeploymentCell(std[i])
	}
	return cells
}

func (d *deployment) fromCells(cells []dataDispose.DataCell) []appsv1.Deployment {
	deployments := make([]appsv1.Deployment, len(cells))
	for i := range cells {
		deployments[i] = appsv1.Deployment(cells[i].(dataDispose.DeploymentCell))
	}

	return deployments
}
