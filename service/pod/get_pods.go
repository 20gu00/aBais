package pod

import (
	"context"
	"errors"
	dataDispose "github.com/20gu00/aBais/common/data_dispose"

	"go.uber.org/zap"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type PodsResp struct {
	Items []corev1.Pod `json:"items"`
	Total int          `json:"total"`
}

func GetPods(client *kubernetes.Clientset, filterName, namespace string, limit, page int) (podsResp *PodsResp, err error) {
	// 获取podList类型的pod列表
	podList, err := client.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		zap.L().Error("获取Pod列表失败, ", zap.Error(err))
		return nil, errors.New("获取Pod列表失败, " + err.Error())
	}
	// 实例化dataSelector对象
	selectableData := &dataDispose.DataSelector{
		GenericDataList: p.toCells(podList.Items),
		dataSelectQuery: &DataSelectQuery{
			FilterQuery: &FilterQuery{Name: filterName},
			PaginateQuery: &PaginateQuery{
				Limit: limit,
				Page:  page,
			},
		},
	}
	//先过滤
	filtered := selectableData.Filter()
	total := len(filtered.GenericDataList)
	//再排序和分页
	data := filtered.Sort().Paginate()

	//将[]DataCell类型的pod列表转为v1.pod列表
	pods := p.fromCells(data.GenericDataList)

	return &PodsResp{
		Items: pods,
		Total: total,
	}, nil
}
