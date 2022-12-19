package namespace

import (
	k8sClient "github.com/20gu00/aBais/common/k8s-clientset"
	"github.com/20gu00/aBais/common/response"
	param "github.com/20gu00/aBais/model/param/ns"
	service "github.com/20gu00/aBais/service/namespace"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 获取namespace详情
func GetNamespaceDetail(ctx *gin.Context) {
	// 1.参数
	params := new(param.GetNsDetailInput)
	if err := ctx.Bind(params); err != nil {
		zap.L().Error("C-GetNamespaceDetail 绑定请求参数失败", zap.Error(err))
		response.RespErr(ctx, response.CodeInvalidParam)
		return
	}

	// 2.service
	client, err := k8sClient.K8s.GetK8sClient(params.Cluster)
	if err != nil {
		zap.L().Error("C-GetNamespaceDetail 获取k8s的client失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeGetK8sClientErr)
		return
	}

	data, err := service.Namespace.GetNamespaceDetail(client, params.NamespaceName)
	if err != nil {
		zap.L().Error("C-GetNamespaceDetail 获取ns详情失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeGetNsDetailErr)
		return
	}

	// 3.resp
	response.RespOK(ctx, "获取Namespace详情成功", data)
}
