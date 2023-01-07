package serviceAccount

import (
	k8sClient "github.com/20gu00/aBais/common/k8s-clientset"
	"github.com/20gu00/aBais/common/response"
	service "github.com/20gu00/aBais/service/service-account"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 创建job
func CreateSa(ctx *gin.Context) {
	var (
		SaCreate = new(service.SaCreate)
		err      error
	)

	// 1.参数
	if err = ctx.ShouldBindJSON(SaCreate); err != nil {
		zap.L().Error("C-CreateSa 绑定请求参数失败", zap.Error(err))
		response.RespErr(ctx, response.CodeInvalidParam)
		return
	}
	//指定的cluster的client
	client, err := k8sClient.K8s.GetK8sClient(SaCreate.Cluster)
	if err != nil {
		zap.L().Error("C-CreateSa 获取k8s的client失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeGetK8sClientErr)
		return
	}
	if err = service.Sa.CreateSa(client, SaCreate); err != nil {
		zap.L().Error("C-CreateSa 创建sa失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeCreateserviceAccountErr) //  前段可以使用
		return
	}

	// 3.resp
	response.RespOK(ctx, "创建sa成功", nil)
}
