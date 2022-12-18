package ingress

import (
	k8sClient "github.com/20gu00/aBais/common/k8s-clientset"
	"github.com/20gu00/aBais/common/response"
	param "github.com/20gu00/aBais/model/param/ingress"
	"github.com/20gu00/aBais/service/ingress"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 删除ingress
func DeleteIngress(ctx *gin.Context) {
	// 1.ingress
	params := new(param.CommonIngressDetailInput)
	// DELETE ShouldBindJSON
	if err := ctx.ShouldBindJSON(params); err != nil {
		zap.L().Error("C-DeleteIngress 绑定请求参数失败", zap.Error(err))
		response.RespErr(ctx, response.CodeInvalidParam)
		return
	}

	// 2.service
	client, err := k8sClient.K8s.GetK8sClient(params.Cluster)
	if err != nil {
		zap.L().Error("C-DeleteIngress 获取k8s的client失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeGetK8sClientErr)
		return
	}

	err = ingress.Ingress.DeleteIngress(client, params.IngressName, params.Namespace)
	if err != nil {
		zap.L().Error("C-DeleteIngress 删除ingress失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeDeleteIngressErr)
		return
	}

	// 3.resp
	response.RespOK(ctx, "删除Ingress成功", nil)
}
