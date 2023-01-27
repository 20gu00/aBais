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
	// 返回文件的指定路径的文件fileInfo
	_, err := os.Stat(filePath)
	if err != nil || os.IsNotExist(err) {
		zap.L().Error("chart文件不存在", zap.Error(err))
		return errors.New("chart文件不存在" + err.Error())
	}
	// 删除文件
	err = os.Remove(filePath)
	if err != nil {
		zap.L().Error("S-DeleteChartFile chart文件删除失败", zap.Error(err))
		return errors.New("S-DeleteChartFile chart文件删除失败" + err.Error())
	}
	return nil
}
