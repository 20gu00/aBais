package service

import (
	"errors"
	"github.com/20gu00/aBais/common/config"
	"go.uber.org/zap"
	"os"
)

//chart文件删除
func (*helmStore) DeleteChartFile(chart string) error {
	filePath := config.Config.UploadPath + "/" + chart
	// not exist,ok
	_, err := os.Stat(filePath)
	if err != nil || os.IsNotExist(err) {
		zap.L().Error("chart文件不存在", zap.Error(err))
		return errors.New("chart文件不存在" + err.Error())
	}
	err = os.Remove(filePath)
	if err != nil {
		zap.L().Error("S-DeleteChartFile chart文件删除失败", zap.Error(err))
		return errors.New("S-DeleteChartFile chart文件删除失败" + err.Error())
	}
	return nil
}
