package deployment

import (
	"fmt"

	k8sClient "github.com/20gu00/aBais/common/k8s-clientset"
	"github.com/20gu00/aBais/common/response"
	param "github.com/20gu00/aBais/model/param/deployment"
	"github.com/20gu00/aBais/service/deployment"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 水平扩缩容
func ScaleDeployment(ctx *gin.Context) {
	// 1.参数
	params := new(param.ScaleDeployment)
	if err := ctx.ShouldBindJSON(params); err != nil {
		zap.L().Error("C-ScaleDeployment 绑定请求参数失败", zap.Error(err))
		response.RespErr(ctx, response.CodeInvalidParam)
		return
	}

	// 2.service
	client, err := k8sClient.K8s.GetK8sClient(params.Cluster)
	if err != nil {
		zap.L().Error("C-ScaleDeployment 获取k8s的client失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeGetK8sClientErr)
		return
	}

	data, err := deployment.Deployment.ScaleDeployment(client, params.DeploymentName, params.Namespace, params.ScaleNum)
	if err != nil {
		zap.L().Error("C-ScaleDeployment 调整deployment副本数目失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeScaleDeploymentErr)
		return
	}

	// 3.resp
	response.RespOK(ctx, "获取Deployment列表成功", fmt.Sprintf("最新副本数: %d", data))
}
