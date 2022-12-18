package secret

import (
	"context"
	"errors"

	dataDispose "github.com/20gu00/aBais/common/data-dispose"

	"go.uber.org/zap"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

var Secret secret

type secret struct{}

type SecretsResp struct {
	Items []corev1.Secret `json:"items"`
	Total int             `json:"total"`
}

// 获取secret列表 过滤、排序、分页
func (s *secret) GetSecrets(client *kubernetes.Clientset, filterName, namespace string, limit, page int) (secretsResp *SecretsResp, err error) {
	// 获取secretList类型的secret列表
	secretList, err := client.CoreV1().Secrets(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		zap.L().Error("S-GetSecrets 获取Secret列表失败", zap.Error(err))
		return nil, errors.New("获取Secret列表失败, " + err.Error())
	}
	// 将secretList中的secret列表(Items)，放进dataselector对象中，进行排序
	selectableData := &dataDispose.DataSelector{
		GenericDataList: s.toCells(secretList.Items),
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
	secrets := s.fromCells(data.GenericDataList)

	return &SecretsResp{
		Items: secrets,
		Total: total,
	}, nil
}

func (s *secret) toCells(std []corev1.Secret) []dataDispose.DataCell {
	cells := make([]dataDispose.DataCell, len(std))
	for i := range std {
		cells[i] = dataDispose.SecretCell(std[i])
	}
	return cells
}

func (s *secret) fromCells(cells []dataDispose.DataCell) []corev1.Secret {
	secrets := make([]corev1.Secret, len(cells))
	for i := range cells {
		secrets[i] = corev1.Secret(cells[i].(dataDispose.SecretCell))
	}

	return secrets
}