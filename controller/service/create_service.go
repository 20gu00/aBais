package service

import (
	k8sClient "github.com/20gu00/aBais/common/k8s-clientset"
	"github.com/20gu00/aBais/common/response"
	k8sSvc "github.com/20gu00/aBais/service/k8s-service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 创建service
func CreateService(ctx *gin.Context) {
	var (
		serviceCreate = new(k8sSvc.ServiceCreate)
		err           error
	)

	// 1.参数
	if err = ctx.ShouldBindJSON(serviceCreate); err != nil {
		zap.L().Error("C-CreateService 绑定请求参数失败", zap.Error(err))
		response.RespErr(ctx, response.CodeInvalidParam)
		return
	}

	// 2.sedrvice
	client, err := k8sClient.K8s.GetK8sClient(serviceCreate.Cluster)
	if err != nil {
		zap.L().Error("C-CreateService 获取k8s的client失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeGetK8sClientErr)
		return
	}

	if err = k8sSvc.K8sService.CreateService(client, serviceCreate); err != nil {
		zap.L().Error("C-CreateService 创建service失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeCreateSvcErr)
	}

	// 3.resp
	response.RespOK(ctx, "创建service成功", nil)
}
