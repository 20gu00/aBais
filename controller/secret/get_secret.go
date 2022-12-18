package secret

import (
	k8sClient "github.com/20gu00/aBais/common/k8s-clientset"
	"github.com/20gu00/aBais/common/response"
	param "github.com/20gu00/aBais/model/param/secret"
	"github.com/20gu00/aBais/service/secret"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 获取secret列表，支持过滤、排序、分页
func GetSecrets(ctx *gin.Context) {
	// 1.参数
	params := new(param.GetSecretInput)
	if err := ctx.Bind(params); err != nil {
		zap.L().Error("C-GetSecrets 绑定请求参数失败", zap.Error(err))
		response.RespErr(ctx, response.CodeInvalidParam)
		return
	}

	// 2.service
	client, err := k8sClient.K8s.GetK8sClient(params.Cluster)
	if err != nil {
		zap.L().Error("C-GetSecrets 获取k8s的client失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeGetK8sClientErr)
		return
	}

	data, err := service.Secret.GetSecrets(client, params.FilterName, params.Namespace, params.Limit, params.Page)
	if err != nil {
		zap.L().Error("C-GetSecrets 创建service失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeGetSecretErr)
		return
	}

	// 3.resp
	response.RespOK(ctx, "获取Secret列表成功", data)
}
