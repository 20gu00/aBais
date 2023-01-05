package HelmStore

import (
	"github.com/20gu00/aBais/common/response"
	"github.com/20gu00/aBais/model"
	service "github.com/20gu00/aBais/service/helm"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// chart新增
func AddChart(ctx *gin.Context) {
	params := new(model.HelmChart)
	if err := ctx.ShouldBindJSON(params); err != nil {
		zap.L().Error("C-AddChart 绑定请求参数失败", zap.Error(err))
		response.RespErr(ctx, response.CodeInvalidParam)
		return
	}
	err := service.HelmStore.AddChart(params)
	if err != nil {
		zap.L().Error("C-AddChart 新增chart失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeAddChartErr)
		return
	}

	response.RespOK(ctx, "新增chart成功", nil)
}
