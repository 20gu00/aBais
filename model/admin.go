package model

import "github.com/jinzhu/gorm"

type Admin struct {
	gorm.Model
	UserID   int64  `grom:"user_id"` // 雪花算法64bit
	Username string `grom:"username"`
	Password string `grom:"password"`
	// 不需要数据库处理
	Token string
}

func (*Admin) TableName() string {
	return "admin"
}
