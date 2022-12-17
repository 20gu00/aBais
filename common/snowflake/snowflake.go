package snowflake

import (
	"time"

	sf "github.com/bwmarrin/snowflake"
)

//每个分布式节点
var node *sf.Node

//开始时间 机器唯一标示id  uid
func InitSnowFlake(startTime string, machineID int64) (err error) {
	var st time.Time
	//string-->time.Time
	st, err = time.Parse("2006-01-02", startTime) //2006-01-02 15:04:05
	if err != nil {
		return
	}
	sf.Epoch = st.UnixNano() / 1000000 //毫秒
	node, err = sf.NewNode(machineID)  //id简单使用可以写死
	return
}
func GenID() int64 {
	return node.Generate().Int64() //也可以是string
}
