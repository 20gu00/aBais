package common

import (
	"encoding/json"
	"go.uber.org/zap"

	"github.com/20gu00/aBais/common/config"
	service "github.com/20gu00/aBais/service/event"
)

func EventWatch() {
	// 初始化
	kubeMap := map[string]string{}
	if err := json.Unmarshal([]byte(config.Config.KubeConfigs), &kubeMap); err != nil {
		zap.L().Error("eventWatch 序列化kubeconfig信息失败", zap.Error(err))
	}
	// 先是kubeMap 左值仅一次创建
	for idx, _ := range kubeMap {
		// 直接用值而不是idx的指针
		go func() {
			service.Event.WatchEventTask(idx)
		}()
	}
}
