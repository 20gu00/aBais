package HelmStore

import (
	"github.com/20gu00/aBais/common/response"
	"github.com/20gu00/aBais/model"
	service "github.com/20gu00/aBais/service/helm"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// chart更新
func UpdateChart(ctx *gin.Context) {
	params := new(model.HelmChart)
	if err := ctx.ShouldBindJSON(params); err != nil {
		zap.L().Error("C-UpdateChart 绑定请求参数失败", zap.Error(err))
		response.RespErr(ctx, response.CodeInvalidParam)
		return
	}
	err := service.HelmStore.UpdateChart(params)
	if err != nil {
		zap.L().Error("C-UpdateChart 更新chart失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeUpdateChartErr)
		return
	}

	response.RespOK(ctx, "更新chart成功", nil)
}