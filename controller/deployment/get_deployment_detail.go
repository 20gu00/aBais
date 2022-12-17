package deployment

import (
	k8sClient "github.com/20gu00/aBais/common/k8s-clientset"
	"github.com/20gu00/aBais/common/response"
	"github.com/20gu00/aBais/service/deployment"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 获取deployment详情
func GetDeploymentDetail(ctx *gin.Context) {
	// 1.参数
	params := new(struct {
		DeploymentName string `form:"deployment_name"`
		Namespace      string `form:"namespace"`
		Cluster        string `form:"cluster"`
	})
	if err := ctx.Bind(params); err != nil {
		zap.L().Error("C-GetDeploymentsDetail 绑定请求参数失败", zap.Error(err))
		response.RespErr(ctx, response.CodeInvalidParam)
		return
	}

	// 2.service
	client, err := k8sClient.K8s.GetK8sClient(params.Cluster)
	if err != nil {
		zap.L().Error("C-GetDeploymentsDetail 获取k8s的client失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeGetK8sClientErr)
		return
	}

	data, err := deployment.Deployment.GetDeploymentDetail(client, params.DeploymentName, params.Namespace)
	if err != nil {
		zap.L().Error("C-GetDeploymentsDetail 获取deployment详情失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeGetDeploymentDetailErr)
		return
	}

	// 3.resp
	response.RespOK(ctx, "获取Deployment详情成功", data)
}
