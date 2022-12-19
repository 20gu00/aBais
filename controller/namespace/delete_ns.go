package namespace

import (
	k8sClient "github.com/20gu00/aBais/common/k8s-clientset"
	"github.com/20gu00/aBais/common/response"
	param "github.com/20gu00/aBais/model/param/ns"
	service "github.com/20gu00/aBais/service/namespace"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 删除namespace
func DeleteNamespace(ctx *gin.Context) {
	// 1.参数
	params := new(param.DeleteNsInput)
	// DELETE ShouldBindJSON
	if err := ctx.ShouldBindJSON(params); err != nil {
		zap.L().Error("C-DeleteNamespace 绑定请求参数失败", zap.Error(err))
		response.RespErr(ctx, response.CodeInvalidParam)
		return
	}

	// service
	client, err := k8sClient.K8s.GetK8sClient(params.Cluster)
	if err != nil {
		zap.L().Error("C-DeleteNamespace 获取k8s的client失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeGetK8sClientErr)
		return
	}

	err = service.Namespace.DeleteNamespace(client, params.NamespaceName)
	if err != nil {
		zap.L().Error("C-DeleteNamespace 删除ns失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeDeleteNsErr)
		return
	}

	// 3.resp
	response.RespOK(ctx, "删除Namespace成功", nil)
}
