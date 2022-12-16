package admin

import (
	"errors"
	"github.com/20gu00/aBais/common/config"
	"go.uber.org/zap"
)

func Login(userName string, password string) error {
	if userName == config.Config.AdminUser && password == config.Config.AdminPassword {
		return nil
	} else {
		zap.L().Error("登录失败, 用户名或密码错误")
		return errors.New("登录失败, 用户名或密码错误")
	}
}
