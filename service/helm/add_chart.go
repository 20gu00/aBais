package service

import (
	"errors"

	"github.com/20gu00/aBais/dao"
	"github.com/20gu00/aBais/model"
)

//chart新增
func (*helmStore) AddChart(chart *model.HelmChart) error {
	_, has, err := dao.Chart.Has(chart.Name)
	if err != nil {
		return err
	}
	if has {
		return errors.New("该chart已存在，请重新添加")
	}
	if err := dao.Chart.Add(chart); err != nil {
		return err
	}
	return nil
}
