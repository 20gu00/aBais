package model

import (
	"time"
)

type HelmChart struct {
	// gorm.Model
	ID        uint `json:"id" gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`

	Name     string `json:"name"`
	FileName string `json:"file_name" gorm:"column: file_name"`
	IconUrl  string `json:"icon_url" gorm:"column: icon_url"`
	Version  string `json:"version"`
	Describe string `json:"describe"`
}

// 默认表名是HelmCharts
func (*HelmChart) TableName() string {
	return "helm_chart"
}
