package HelmStore

import (
	helmClient "github.com/20gu00/aBais/common/helm-client"
	"github.com/20gu00/aBais/common/response"
	param "github.com/20gu00/aBais/model/param/release"
	service "github.com/20gu00/aBais/service/helm"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// release详情
func DetailRelease(ctx *gin.Context) {
	params := new(param.DetailReleaseInput)
	if err := ctx.Bind(params); err != nil {
		zap.L().Error("C-DetailRelease 绑定请求参数失败", zap.Error(err))
		response.RespErr(ctx, response.CodeInvalidParam)
		return
	}
	actionConfig, err := helmClient.HelmConfig.GetAc(params.Cluster, params.Namespace)
	if err != nil {
		zap.L().Error("C-DetailRelease 获取action的config失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeActionConfigErr)
		return
	}
	data, err := service.HelmStore.DetailRelease(actionConfig, params.Release)
	if err != nil {
		zap.L().Error("C-DetailRelease 获取release详情失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeDetailReleaseErr)
		return
	}

	response.RespOK(ctx, "获取release详情成功", data)
}
