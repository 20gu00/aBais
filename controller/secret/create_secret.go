package secret

import (
	k8sClient "github.com/20gu00/aBais/common/k8s-clientset"
	"github.com/20gu00/aBais/common/response"
	service "github.com/20gu00/aBais/service/secret"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreateSecret(ctx *gin.Context) {
	var (
		secretCreate = new(service.SecretCreate)
		err          error
	)

	// 1.参数
	if err = ctx.ShouldBindJSON(secretCreate); err != nil {
		zap.L().Error("C-CreateSecret 绑定请求参数失败", zap.Error(err))
		response.RespErr(ctx, response.CodeInvalidParam)
		return
	}
	//指定的cluster的client
	client, err := k8sClient.K8s.GetK8sClient(secretCreate.Cluster)
	if err != nil {
		zap.L().Error("C-CreateSecret 获取k8s的client失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeGetK8sClientErr)
		return
	}
	if err = service.Secret.CreateSecret(client, secretCreate); err != nil {
		zap.L().Error("C-CreateSecret 创建secret失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeCreateSecretErr) //  前段可以使用
		return
	}

	// 3.resp
	response.RespOK(ctx, "创建Secret成功", nil)
}
