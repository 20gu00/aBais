package cm

import (
	"context"
	"errors"

	dataDispose "github.com/20gu00/aBais/common/data-dispose"

	"go.uber.org/zap"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

var ConfigMap configMap

type configMap struct{}

type ConfigMapsResp struct {
	Items []corev1.ConfigMap `json:"items"`
	Total int                `json:"total"`
}

// 获取configmap列表 过滤、排序、分页
func (c *configMap) GetConfigMaps(client *kubernetes.Clientset, filterName, namespace string, limit, page int) (configMapsResp *ConfigMapsResp, err error) {
	// 获取configMapList类型的configMap列表
	configMapList, err := client.CoreV1().ConfigMaps(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		zap.L().Error("S-GetConfigMaps 获取ConfigMap列表失败", zap.Error(err))
		return nil, errors.New("获取ConfigMap列表失败, " + err.Error())
	}
	// 将configMapList中的configMap列表(Items)，放进dataselector对象中，进行排序
	selectableData := &dataDispose.DataSelector{
		GenericDataList: c.toCells(configMapList.Items),
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

	// 将[]DataCell类型的configmap列表转为v1.configmap列表
	configMaps := c.fromCells(data.GenericDataList)

	return &ConfigMapsResp{
		Items: configMaps,
		Total: total,
	}, nil
}

func (c *configMap) toCells(std []corev1.ConfigMap) []dataDispose.DataCell {
	cells := make([]dataDispose.DataCell, len(std))
	for i := range std {
		cells[i] = dataDispose.ConfigMapCell(std[i])
	}
	return cells
}

func (c *configMap) fromCells(cells []dataDispose.DataCell) []corev1.ConfigMap {
	configMaps := make([]corev1.ConfigMap, len(cells))
	for i := range cells {
		configMaps[i] = corev1.ConfigMap(cells[i].(dataDispose.ConfigMapCell))
	}

	return configMaps
}