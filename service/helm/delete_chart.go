package service

import (
	"fmt"

	"github.com/20gu00/aBais/dao"
	"github.com/20gu00/aBais/model"
)

//Chart更新
func (h *helmStore) UpdateChart(chart *model.HelmChart) error {
	oldChart, _, err := dao.Chart.Has(chart.Name)
	if err != nil {
		return err
	}
	fmt.Println(chart.FileName, oldChart.FileName)
	if chart.FileName != "" && chart.FileName != oldChart.FileName {
		err = h.DeleteChartFile(oldChart.FileName)
		if err != nil {
			return err
		}
	}
	return dao.Chart.Update(chart)
}

//Chart删除
func (h *helmStore) DeleteChart(chart *model.HelmChart) error {
	//删除文件
	err := h.DeleteChartFile(chart.FileName)
	if err != nil {
		return err
	}
	//删除数据
	return dao.Chart.Delete(chart.ID)
}