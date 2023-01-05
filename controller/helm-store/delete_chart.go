package HelmStore

import (
	"github.com/20gu00/aBais/common/response"
	"github.com/20gu00/aBais/model"
	service "github.com/20gu00/aBais/service/helm"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// chart删除
func DeleteChart(ctx *gin.Context) {
	params := new(model.HelmChart)
	if err := ctx.ShouldBindJSON(params); err != nil {
		zap.L().Error("C-DeleteChart 绑定请求参数失败", zap.Error(err))
		response.RespErr(ctx, response.CodeInvalidParam)
		return
	}
	err := service.HelmStore.DeleteChart(params)
	if err != nil {
		zap.L().Error("C-DeleteChart 删除chart失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeDeleteChartErr)
		return
	}

	response.RespOK(ctx, "删除chart成功", nil)
}
