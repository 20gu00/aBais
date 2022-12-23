package service

import "github.com/20gu00/aBais/dao"

var Event event

type event struct{}

// 从数据库中获取event列表
func (*event) GetList(name, cluster string, page, limit int) (*dao.Events, error) {
	data, err := dao.Event.GetList(name, cluster, page, limit)
	if err != nil {
		return nil, err
	}
	return data, nil
}
