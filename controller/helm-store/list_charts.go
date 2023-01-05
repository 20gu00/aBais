package HelmStore

import (
	"github.com/20gu00/aBais/common/response"
	param "github.com/20gu00/aBais/model/param/release"
	service "github.com/20gu00/aBais/service/helm"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// chart列表
func ListCharts(ctx *gin.Context) {
	params := new(param.ListChartInput)
	if err := ctx.Bind(params); err != nil {
		zap.L().Error("C-ListCharts 绑定请求参数失败", zap.Error(err))
		response.RespErr(ctx, response.CodeInvalidParam)
		return
	}
	data, err := service.HelmStore.ListCharts(params.Name, params.Page, params.Limit)
	if err != nil {
		zap.L().Error("C-ListCharts 获取chart列表失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeListChartErr)
		return
	}

	response.RespOK(ctx, "安装release成功", data)
}
