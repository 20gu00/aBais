package HelmStore

import (
	helmClient "github.com/20gu00/aBais/common/helm-client"
	"github.com/20gu00/aBais/common/response"
	param "github.com/20gu00/aBais/model/param/release"
	service "github.com/20gu00/aBais/service/helm"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// release卸载
func UninstallRelease(ctx *gin.Context) {
	params := new(param.UninstallReleaseInput)
	if err := ctx.ShouldBindJSON(params); err != nil {
		zap.L().Error("C-UninstallRelease 绑定请求参数失败", zap.Error(err))
		response.RespErr(ctx, response.CodeInvalidParam)
		return
	}
	actionConfig, err := helmClient.HelmConfig.GetAc(params.Cluster, params.Namespace)
	if err != nil {
		zap.L().Error("C-UninstallRelease 获取action的config失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeActionConfigErr)
		return
	}
	err = service.HelmStore.UninstallRelease(actionConfig, params.Release, params.Namespace)
	if err != nil {
		zap.L().Error("C-UninstallRelease 卸载release失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeUninstallReleaseErr)
		return
	}

	response.RespOK(ctx, "卸载release详情成功", nil)
}
