package service

import (
	"context"
	"errors"

	dataDispose "github.com/20gu00/aBais/common/data-dispose"

	"go.uber.org/zap"
	batchv1 "k8s.io/api/batch/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

var Job job

type job struct{}

type JobsResp struct {
	Items []batchv1.Job `json:"items"`
	Total int           `json:"total"`
}

func (j *job) GetJobs(client *kubernetes.Clientset, filterName, namespace string, limit, page int) (deploymentsResp *JobsResp, err error) {
	jobList, err := client.BatchV1().Jobs(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		zap.L().Error("S-GetJobs 获取Job列表失败", zap.Error(err))
		return nil, errors.New("获取Job列表失败, " + err.Error())
	}

	selectableData := &dataDispose.DataSelector{
		GenericDataList: j.toCells(jobList.Items),
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

	jobs := j.fromCells(data.GenericDataList)

	return &JobsResp{
		Items: jobs,
		Total: total,
	}, nil
}

func (j *job) toCells(std []batchv1.Job) []dataDispose.DataCell {
	cells := make([]dataDispose.DataCell, len(std))
	for i := range std {
		cells[i] = dataDispose.JobCell(std[i])
	}
	return cells
}

func (j *job) fromCells(cells []dataDispose.DataCell) []batchv1.Job {
	jobs := make([]batchv1.Job, len(cells))
	for i := range cells {
		jobs[i] = batchv1.Job(cells[i].(dataDispose.JobCell))
	}

	return jobs
}
