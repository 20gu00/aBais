package HelmStore

import (
	"github.com/20gu00/aBais/common/response"
	param "github.com/20gu00/aBais/model/param/release"
	service "github.com/20gu00/aBais/service/helm"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// chart文件删除
func DeleteChartFile(ctx *gin.Context) {
	params := new(param.DeleteChartFileInput)
	if err := ctx.ShouldBindJSON(params); err != nil {
		zap.L().Error("C-DeleteChartFile 绑定请求参数失败", zap.Error(err))
		response.RespErr(ctx, response.CodeInvalidParam)
		return
	}
	err := service.HelmStore.DeleteChartFile(params.Chart)
	if err != nil {
		zap.L().Error("C-DeleteChartFile 删除chart file失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeDeleteChartFileErr)
		return
	}

	response.RespOK(ctx, "删除chart file成功", nil)
}
