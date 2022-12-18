package ingress

import (
	k8sClient "github.com/20gu00/aBais/common/k8s-clientset"
	"github.com/20gu00/aBais/common/response"
	"github.com/20gu00/aBais/service/ingress"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 创建ingress
func CreateIngress(ctx *gin.Context) {
	var (
		ingressCreate = new(ingress.IngressCreate)
		err           error
	)

	// 1.参数
	if err = ctx.ShouldBindJSON(ingressCreate); err != nil {
		zap.L().Error("C-CreateIngress 绑定请求参数失败", zap.Error(err))
		response.RespErr(ctx, response.CodeInvalidParam)
		return
	}

	// 2.service
	client, err := k8sClient.K8s.GetK8sClient(ingressCreate.Cluster)
	if err != nil {
		zap.L().Error("C-CreateIngress 获取k8s的client失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeGetK8sClientErr)
		return
	}

	if err = ingress.Ingress.CreateIngress(client, ingressCreate); err != nil {
		zap.L().Error("C-CreateIngress 创建ingress失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeCreateIngressErr)
		return
	}

	// 3.resp
	response.RespOK(ctx, "创建Ingress成功", nil)
}
