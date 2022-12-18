package secret

import (
	k8sClient "github.com/20gu00/aBais/common/k8s-clientset"
	"github.com/20gu00/aBais/common/response"
	param "github.com/20gu00/aBais/model/param/secret"
	"github.com/20gu00/aBais/service/secret"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 删除secret
func DeleteSecret(ctx *gin.Context) {
	// 1.参数
	params := new(param.DeleteSecretInput)
	// DELETE ShouldBindJSON
	if err := ctx.ShouldBindJSON(params); err != nil {
		zap.L().Error("C-DeleteSecret 绑定请求参数失败", zap.Error(err))
		response.RespErr(ctx, response.CodeInvalidParam)
		return
	}

	// 2.service
	client, err := k8sClient.K8s.GetK8sClient(params.Cluster)
	if err != nil {
		zap.L().Error("C-DeleteSecret 获取k8s的client失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeGetK8sClientErr)
		return
	}

	err = secret.Secret.DeleteSecret(client, params.SecretName, params.Namespace)
	if err != nil {
		zap.L().Error("C-DeleteSecret 删除secret失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeDeleteSecretErr)
		return
	}

	// 3.resp
	response.RespOK(ctx, "删除Secret成功", nil)
}
