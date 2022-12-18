package deployment

import (
	k8sClient "github.com/20gu00/aBais/common/k8s-clientset"
	"github.com/20gu00/aBais/common/response"
	"github.com/20gu00/aBais/service/deployment"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 创建deployment
func CreateDeployment(ctx *gin.Context) {
	var (
		deployCreate = new(service.DeployCreate)
		err          error
	)

	// 1.参数
	if err = ctx.ShouldBindJSON(deployCreate); err != nil {
		zap.L().Error("C-CreateDeployment 绑定请求参数失败", zap.Error(err))
		response.RespErr(ctx, response.CodeInvalidParam)
		return
	}
	client, err := k8sClient.K8s.GetK8sClient(deployCreate.Cluster)
	if err != nil {
		zap.L().Error("C-CreateDeployment 获取k8s的client失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeGetK8sClientErr)
		return
	}
	if err = service.Deployment.CreateDeployment(client, deployCreate); err != nil {
		zap.L().Error("C-CreateDeployment 创建deployment失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeDeleteDeploymentErr)
		return
	}

	// 3.resp
	response.RespOK(ctx, "创建Deployment成功", nil)
}
