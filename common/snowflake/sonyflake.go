package snowflake

import (
	"errors"
	"fmt"
	"github.com/sony/sonyflake"
	"time"
)

var (
	sonyFlake     *sonyflake.Sonyflake
	sonyMachineID uint16 //16bit
)

func getMachineID() (uint16, error) {
	return sonyMachineID, nil
}

func InitSonyFlake(startTime string, machineId uint16) (err error) {
	var st time.Time
	st, err = time.Parse("2006-01-02", startTime)
	if err != nil {
		return err
	}

	settings := sonyflake.Settings{
		StartTime: st,
		MachineID: getMachineID, //函数
	}
	sonyFlake = sonyflake.NewSonyflake(settings)
	return nil
}

//生成id
//64bit
func GetID() (id uint64, err error) {
	if sonyFlake == nil {
		//fmt.Errorf
		err = fmt.Errorf("sony雪花算法初始化sonyflake失败", err)
		return //直接return
	}

	id, err = sonyFlake.NextID()
	if err != nil {
		err = errors.New("获取生成的id失败")
		return
	}
	return id, nil
}
