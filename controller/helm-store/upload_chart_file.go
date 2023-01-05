package HelmStore

import (
	"github.com/20gu00/aBais/common/response"
	service "github.com/20gu00/aBais/service/helm"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// chart文件上传
func UploadChartFile(ctx *gin.Context) {
	file, header, err := ctx.Request.FormFile("chart")
	if err != nil {
		zap.L().Error("C-UploadChartFile 获取上传信息失败", zap.Error(err))
		response.RespErr(ctx, response.CodeGetUploadMessageErr)
		return
	}
	err = service.HelmStore.UploadChartFile(file, header)
	if err != nil {
		zap.L().Error("C-UploadChartFile 上传chart file失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeUploadChartFileErr)
		return
	}

	response.RespOK(ctx, "上传chart file成功", nil)
}
