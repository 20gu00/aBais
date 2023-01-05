package service

import (
	"errors"

	"go.uber.org/zap"
	"helm.sh/helm/v3/pkg/action"
)

//release卸载
func (*helmStore) UninstallRelease(actionConfig *action.Configuration, release, namespace string) error {
	client := action.NewUninstall(actionConfig)
	_, err := client.Run(release)
	if err != nil {
		zap.L().Error("卸载release失败", zap.Error(err))
		return errors.New("卸载release失败" + err.Error())
	}
	return nil
}
