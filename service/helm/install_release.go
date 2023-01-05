package service

import (
	"errors"
	"strings"

	"github.com/20gu00/aBais/common/config"

	"go.uber.org/zap"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chart/loader"
)

//release安装
func (*helmStore) InstallRelease(actionConfig *action.Configuration, release, chart, namespace string) error {
	client := action.NewInstall(actionConfig)
	client.ReleaseName = release
	client.Namespace = namespace

	splitChart := strings.Split(chart, ".")
	if splitChart[len(splitChart)-1] == "tgz" && !strings.Contains(chart, ":") {
		chart = config.Config.UploadPath + "/" + chart
	}
	chartRequested, err := loader.Load(chart)
	if err != nil {
		zap.L().Error("加载chart文件失败" + err.Error())
		return errors.New("加载chart文件失败" + err.Error())
	}
	vals := map[string]interface{}{}
	_, err = client.Run(chartRequested, vals)
	if err != nil {
		zap.L().Error("S-InstallRelease 安装release失败", zap.Error(err))
		return errors.New("S-InstallRelease 安装release失败" + err.Error())
	}
	return nil
}
