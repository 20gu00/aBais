package dao

import (
	"errors"
	"fmt"
	"go.uber.org/zap"
	"time"

	"github.com/20gu00/aBais/dao/db"
	"github.com/20gu00/aBais/model"

	"github.com/jinzhu/gorm"
)

// 直接暴露(也可以做成interface)
var Event event

type event struct{}

type Events struct {
	Items []*model.K8sEvent `json:"items"`
	Total int               `json:"total"`
}

func (*event) GetList(name, cluster string, page, limit int) (*Events, error) {
	// 定义分页数据的起始位置
	startSet := (page - 1) * limit

	// 定义数据库查询返回内容
	var (
		eventList []*model.K8sEvent
		total     int
	)

	// 数据库查询，Limit方法用于限制条数a，Offset方法设置起始位置b limit(a,b)
	tx := db.GORM.
		Model(&model.K8sEvent{}).
		Where("name like ? and cluster = ?", "%"+name+"%", cluster).
		Count(&total).
		Limit(limit).
		Offset(startSet).
		Order("id desc").
		Find(&eventList)

	if tx.Error != nil {
		zap.L().Error("获取Event列表失败", zap.Error(tx.Error))
		return nil, errors.New(fmt.Sprintf("获取Event列表失败,%v\n", tx.Error))
	}

	return &Events{
		Items: eventList,
		Total: total,
	}, nil
}

// 新增event
func (*event) Add(event *model.K8sEvent) error {
	tx := db.GORM.Create(&event)
	if tx.Error != nil {
		zap.L().Error("添加Event失败", zap.Error(tx.Error))
		return errors.New(fmt.Sprintf("添加Event失败, %v\n", tx.Error))
	}
	return nil
}

// 查询单个event
func (*event) HasEvent(name, kind, namespace, reason string, eventTime time.Time, cluster string) (*model.K8sEvent, bool, error) {
	data := &model.K8sEvent{}
	tx := db.GORM.Where("name = ? and kind = ? and namespace = ? and reason = ? and event_time = ? and cluster = ?",
		name, kind, namespace, reason, eventTime, cluster).First(&data)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return nil, false, nil
	}
	if tx.Error != nil {
		zap.L().Error("查询Event失败", zap.Error(tx.Error))
		return nil, false, errors.New(fmt.Sprintf("查询Event失败, %v\n", tx.Error))
	}

	return data, true, nil
}
