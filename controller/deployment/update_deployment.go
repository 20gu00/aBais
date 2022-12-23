package deployment

import (
	k8sClient "github.com/20gu00/aBais/common/k8s-clientset"
	"github.com/20gu00/aBais/common/response"
	param "github.com/20gu00/aBais/model/param/deployment"
	"github.com/20gu00/aBais/service/deployment"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 更新deployment
func UpdateDeployment(ctx *gin.Context) {
	// 1.参数
	params := new(param.UpdateDeployment)
	// PUT ShouldBindJSON
	if err := ctx.ShouldBindJSON(params); err != nil {
		zap.L().Error("C-UpdateDeployment 绑定请求参数失败", zap.Error(err))
		response.RespErr(ctx, response.CodeInvalidParam)
		return
	}

	// 2.service
	client, err := k8sClient.K8s.GetK8sClient(params.Cluster)
	if err != nil {
		zap.L().Error("C-UpdateDeployment 获取k8s的client失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeGetK8sClientErr)
		return
	}

	err = service.Deployment.UpdateDeployment(client, params.Namespace, params.Content)
	if err != nil {
		zap.L().Error("C-UpdateDeployment 更新deployment失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeUpdateDeploymentErr)
		return
	}

	// 3.resp
	response.RespOK(ctx, "更新Deployment成功", nil)
}
