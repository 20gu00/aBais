package service

import (
	k8sClient "github.com/20gu00/aBais/common/k8s-clientset"
	"github.com/20gu00/aBais/common/response"
	param "github.com/20gu00/aBais/model/param/service"
	k8sSvc "github.com/20gu00/aBais/service/k8s-service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 获取service详情
func GetServiceDetail(ctx *gin.Context) {
	// 1.参数
	params := new(param.GetServiceDetailInput)
	if err := ctx.Bind(params); err != nil {
		zap.L().Error("C-GetServiceDetail 绑定请求参数失败", zap.Error(err))
		response.RespErr(ctx, response.CodeInvalidParam)
		return
	}

	// 2.service
	client, err := k8sClient.K8s.GetK8sClient(params.Cluster)
	if err != nil {
		zap.L().Error("C-GetServiceDetail 获取k8s的client失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeGetK8sClientErr)
		return
	}

	data, err := k8sSvc.K8sService.GetServicetDetail(client, params.ServiceName, params.Namespace)
	if err != nil {
		zap.L().Error("C-GetServiceDetail 获取service详情失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeGetSvcDetailErr)
		return
	}

	// 3.resp
	response.RespOK(ctx, "获取service详情成功", data)
}