package service

import (
	"context"
	"errors"

	dataDispose "github.com/20gu00/aBais/common/data-dispose"

	"go.uber.org/zap"
	batchv1beta1 "k8s.io/api/batch/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

var CronJob cronjob

type cronjob struct{}

type CronJobsResp struct {
	Items []batchv1beta1.CronJob `json:"items"`
	Total int                    `json:"total"`
}

func (c *cronjob) GetCronJobs(client *kubernetes.Clientset, filterName, namespace string, limit, page int) (deploymentsResp *CronJobsResp, err error) {
	cronJobList, err := client.BatchV1beta1().CronJobs(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		zap.L().Error("S-GetJobs 获取CronJob列表失败", zap.Error(err))
		return nil, errors.New("获取CronJob列表失败, " + err.Error())
	}

	selectableData := &dataDispose.DataSelector{
		GenericDataList: c.toCells(cronJobList.Items),
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

	cronJobs := c.fromCells(data.GenericDataList)

	return &CronJobsResp{
		Items: cronJobs,
		Total: total,
	}, nil
}

func (j *cronjob) toCells(std []batchv1beta1.CronJob) []dataDispose.DataCell {
	cells := make([]dataDispose.DataCell, len(std))
	for i := range std {
		cells[i] = dataDispose.CronJobCell(std[i])
	}
	return cells
}

func (j *cronjob) fromCells(cells []dataDispose.DataCell) []batchv1beta1.CronJob {
	cronJobs := make([]batchv1beta1.CronJob, len(cells))
	for i := range cells {
		cronJobs[i] = batchv1beta1.CronJob(cells[i].(dataDispose.CronJobCell))
	}

	return cronJobs
}
