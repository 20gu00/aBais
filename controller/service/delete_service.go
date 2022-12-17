package service

import (
	k8sClient "github.com/20gu00/aBais/common/k8s-clientset"
	"github.com/20gu00/aBais/common/response"
	param "github.com/20gu00/aBais/model/param/service"
	k8sSvc "github.com/20gu00/aBais/service/k8s-service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 删除service
func DeleteService(ctx *gin.Context) {
	// 1.参数
	params := new(param.DeleteServiceInput)
	// DELETE ShouldBindJSON
	if err := ctx.ShouldBindJSON(params); err != nil {
		zap.L().Error("C-DeleteService 绑定请求参数失败", zap.Error(err))
		response.RespErr(ctx, response.CodeInvalidParam)
		return
	}

	// 2.service
	client, err := k8sClient.K8s.GetK8sClient(params.Cluster)
	if err != nil {
		zap.L().Error("C-DeleteService 获取k8s的client失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeGetK8sClientErr)
		return
	}

	err = k8sSvc.K8sService.DeleteService(client, params.ServiceName, params.Namespace)
	if err != nil {
		zap.L().Error("C-DeleteService 删除service失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeDeleteSvcErr)
		return
	}

	// 3.resp
	response.RespOK(ctx, "删除service成功", nil)
}
