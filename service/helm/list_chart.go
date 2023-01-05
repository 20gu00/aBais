package service

import "github.com/20gu00/aBais/dao"

//chart列表
func (*helmStore) ListCharts(name string, page, limit int) (*dao.Charts, error) {
	return dao.Chart.GetList(name, page, limit)
}
