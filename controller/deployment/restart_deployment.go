package deployment

import (
	k8sClient "github.com/20gu00/aBais/common/k8s-clientset"
	"github.com/20gu00/aBais/common/response"
	param "github.com/20gu00/aBais/model/param/deployment"
	"github.com/20gu00/aBais/service/deployment"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 重启deployment
func RestartDeployment(ctx *gin.Context) {
	// 1.参数
	params := new(param.RestartDeployment)
	// PUT ShouldBindJSON
	if err := ctx.ShouldBindJSON(params); err != nil {
		zap.L().Error("C-RestartDeployment 绑定请求参数失败, ", zap.Error(err))
		response.RespErr(ctx, response.CodeInvalidParam)
		return
	}

	// 2.service
	client, err := k8sClient.K8s.GetK8sClient(params.Cluster)
	if err != nil {
		zap.L().Error("C-RestartDeployment 获取k8s的client失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeGetK8sClientErr)
		return
	}

	err = deployment.Deployment.RestartDeployment(client, params.DeploymentName, params.Namespace)
	if err != nil {
		zap.L().Error("C-RestartDeployment 重启deployment失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeRestartDeploymentErr)
		return
	}

	// 3.resp
	response.RespOK(ctx, "重启Deployment成功", nil)
}
