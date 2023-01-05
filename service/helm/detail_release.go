package service

import (
	"errors"
	"go.uber.org/zap"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/release"
)

//release详情
func (*helmStore) DetailRelease(actionConfig *action.Configuration, release string) (*release.Release, error) {
	client := action.NewGet(actionConfig)
	data, err := client.Run(release)
	if err != nil {
		zap.L().Error("S-DetailRelease 获取release detail失败", zap.Error(err))
		return nil, errors.New("S-DetailRelease 获取release detail失败" + err.Error())
	}
	return data, nil
}
