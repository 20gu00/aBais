package HelmStore

import (
	helmClient "github.com/20gu00/aBais/common/helm-client"
	"github.com/20gu00/aBais/common/response"
	param "github.com/20gu00/aBais/model/param/release"
	service "github.com/20gu00/aBais/service/helm"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// release安装
func InstallRelease(ctx *gin.Context) {
	params := new(param.InstallReleaseInput)
	if err := ctx.ShouldBindJSON(params); err != nil {
		zap.L().Error("C-DetailRelease 绑定请求参数失败", zap.Error(err))
		response.RespErr(ctx, response.CodeInvalidParam)
		return
	}
	actionConfig, err := helmClient.HelmConfig.GetAc(params.Cluster, params.Namespace)
	if err != nil {
		zap.L().Error("C-InstallRelease 获取action的config失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeActionConfigErr)
		return
	}
	err = service.HelmStore.InstallRelease(actionConfig, params.Release, params.Chart, params.Namespace)
	if err != nil {
		zap.L().Error("C-InstallRelease 安装release失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeInstallReleaseErr)
		return
	}

	response.RespOK(ctx, "安装release成功", nil)
}
