package initDo

import (
	"flag"
	"fmt"

	"github.com/20gu00/aBais/common/config"
	"github.com/20gu00/aBais/dao/db"
	"github.com/20gu00/aBais/router"
)

func InitDo() {
	// 加载配置文件需要时间,如果这里是使用goroutine,很可能因为配置文件为加载完成从而获取空值,适当阻塞一下
	var confFile string
	flag.StringVar(&confFile, "conf", "", "配置文件")
	flag.Parse()

	if err := config.ConfigRead(confFile); err != nil {
		fmt.Printf("读取配置文件失败, err:%v\n", err)
		panic(err)
	}

	// 初始化数据库连接
	db.InitDB()

	// 初始化路由配置
	router.InitRouter()
}
