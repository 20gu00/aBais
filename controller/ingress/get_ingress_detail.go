package ingress

import (
	k8sClient "github.com/20gu00/aBais/common/k8s-clientset"
	"github.com/20gu00/aBais/common/response"
	param "github.com/20gu00/aBais/model/param/ingress"
	"github.com/20gu00/aBais/service/ingress"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 获取ingress详情
func GetIngressDetail(ctx *gin.Context) {
	// 1.参数
	params := new(param.CommonIngressDetailInput)
	if err := ctx.Bind(params); err != nil {
		zap.L().Error("C-GetIngressDetail 绑定请求参数失败", zap.Error(err))
		response.RespErr(ctx, response.CodeInvalidParam)
		return
	}

	// 2.service
	client, err := k8sClient.K8s.GetK8sClient(params.Cluster)
	if err != nil {
		zap.L().Error("C-GetIngressDetail 获取k8s的client失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeGetK8sClientErr)
		return
	}

	data, err := ingress.Ingress.GetIngresstDetail(client, params.IngressName, params.Namespace)
	if err != nil {
		zap.L().Error("C-GetIngressDetail 获取ingress详情失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeGetIngressDetailErr)
		return
	}

	// 3.resp
	response.RespOK(ctx, "获取Ingress详情成功", data)
}
