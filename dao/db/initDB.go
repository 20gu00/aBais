package db

import (
	"fmt"
	"time"

	"github.com/20gu00/aBais/common/config"
	"github.com/20gu00/aBais/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go.uber.org/zap"
)

var (
	isInit bool
	// 跨包调用,也可以小写只在model层使用
	GORM *gorm.DB
	err  error
)

// 初始化 连接数据库
func InitDB() {
	// 数据库是否已经初始化了
	if isInit {
		return
	}

	// 数据库配置
	// parseTime自动处理时间  数据库的datetime timestamp <=> go的time.Time  (createdAt)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		config.Config.DBConf.DBUser,
		config.Config.DBConf.DBPassword,
		config.Config.DBConf.DBHost,
		config.Config.DBConf.DBPort,
		config.Config.DBConf.DBName)

	GORM, err = gorm.Open(config.Config.DBType, dsn)
	if err != nil {
		panic("数据库连接失败" + err.Error())
	}

	// 设置日志模式 “true”表示详细日志，“false”表示无日志，默认情况下，只打印错误日志
	GORM.LogMode(config.Config.LogMode)

	// 连接池设置
	// 连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于20，超过的连接会被连接池关闭
	GORM.DB().SetMaxIdleConns(config.Config.DBConf.MaxIdleConns)
	// 设置了连接可复用的最大时间
	GORM.DB().SetMaxOpenConns(config.Config.DBConf.MaxOpenConns)
	// 设置了连接可复用的最大时间
	// 默认单位纳秒
	GORM.DB().SetConnMaxLifetime(time.Duration(config.Config.DBConf.MaxLifeTime) * time.Second)

	// 设置为已经初始化的数据库连接
	isInit = true
	// 自动建表(生产环境不要这么用,ddl都应该交给dba或者运维人员)
	GORM.AutoMigrate(model.HelmChart{}, model.K8sEvent{})
	zap.L().Info("连接数据库成功!")
}

// 关闭
func DBClose() error {
	zap.L().Info("关闭数据库连接")
	return GORM.Close()
}
