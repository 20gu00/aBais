package HelmStore

import (
	helmClient "github.com/20gu00/aBais/common/helm-client"
	"github.com/20gu00/aBais/common/response"
	param "github.com/20gu00/aBais/model/param/release"
	service "github.com/20gu00/aBais/service/helm"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

//已安装的release列表
func ListReleases(ctx *gin.Context) {
	params := new(param.ListReleaseInput)
	if err := ctx.Bind(params); err != nil {
		zap.L().Error("C-ListReleases 绑定请求参数失败", zap.Error(err))
		response.RespErr(ctx, response.CodeInvalidParam)
		return
	}
	actionConfig, err := helmClient.HelmConfig.GetAc(params.Cluster, params.Namespace)
	if err != nil {
		zap.L().Error("C-ListReleases 获取action的config失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeActionConfigErr)
		return
	}
	data, err := service.HelmStore.ListReleases(actionConfig, params.FilterName)
	if err != nil {
		zap.L().Error("C-ListReleases 获取release列表失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeListReleaseErr)
		return
	}

	response.RespOK(ctx, "获取release列表成功", data)
}
