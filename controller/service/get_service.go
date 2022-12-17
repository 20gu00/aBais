package service

import (
	k8sClient "github.com/20gu00/aBais/common/k8s-clientset"
	"github.com/20gu00/aBais/common/response"
	param "github.com/20gu00/aBais/model/param/service"
	k8sSvc "github.com/20gu00/aBais/service/k8s-service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 获取service列表过滤、排序、分页
func GetServices(ctx *gin.Context) {
	// 1.参数
	params := new(param.GetServiceInput)
	if err := ctx.Bind(params); err != nil {
		zap.L().Error("C-GetServices 绑定请求参数失败", zap.Error(err))
		response.RespErr(ctx, response.CodeInvalidParam)
		return
	}

	// 2.service
	client, err := k8sClient.K8s.GetK8sClient(params.Cluster)
	if err != nil {
		zap.L().Error("C-GetServices 获取k8s的client失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeGetK8sClientErr)
		return
	}
	data, err := k8sSvc.K8sService.GetServices(client, params.FilterName, params.Namespace, params.Limit, params.Page)
	if err != nil {
		zap.L().Error("C-GetServices 获取service失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeGetSvcErr)
		return
	}

	// 3.resp
	response.RespOK(ctx, "获取service成功", data)
}
