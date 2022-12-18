package pv

import (
	k8sClient "github.com/20gu00/aBais/common/k8s-clientset"
	"github.com/20gu00/aBais/common/response"
	param "github.com/20gu00/aBais/model/param/pv"
	"github.com/20gu00/aBais/service/pv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 获取pv详情
func GetPvDetail(ctx *gin.Context) {
	// 1.参数
	params := new(param.GetPvDetailInput)
	if err := ctx.Bind(params); err != nil {
		zap.L().Error("C-GetPvDetail 绑定请求参数失败", zap.Error(err))
		response.RespErr(ctx, response.CodeInvalidParam)
		return
	}

	// 2.service
	client, err := k8sClient.K8s.GetK8sClient(params.Cluster)
	if err != nil {
		zap.L().Error("C-GetPvDetail 获取k8s的client失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeGetK8sClientErr)
		return
	}

	data, err := service.Pv.GetPvDetail(client, params.PvName)
	if err != nil {
		zap.L().Error("C-GetPvDetail 获取pv详情失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeGetPvDetailErr)
		return
	}

	// 3.resp
	response.RespOK(ctx, "获取Pv详情成功", data)
}
