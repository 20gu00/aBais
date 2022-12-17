package deployment

import (
	k8sClient "github.com/20gu00/aBais/common/k8s-clientset"
	"github.com/20gu00/aBais/common/response"
	param "github.com/20gu00/aBais/model/param/deployment"
	"github.com/20gu00/aBais/service/deployment"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 获取每个namespace的deployment数量
func GetDeployNumPerNs(ctx *gin.Context) {
	// 1.参数
	params := new(param.GetDeployNumPerNs)
	if err := ctx.Bind(params); err != nil {
		zap.L().Error("C-GetDeployNumPerNs 绑定请求参数失败", zap.Error(err))
		response.RespErr(ctx, response.CodeInvalidParam)
		return
	}

	// 2.serivce
	client, err := k8sClient.K8s.GetK8sClient(params.Cluster)
	if err != nil {
		zap.L().Error("C-GetDeployNumPerNs 获取k8s的client失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeGetK8sClientErr)
		return
	}
	data, err := deployment.Deployment.GetDeployNumPerNs(client)
	if err != nil {
		zap.L().Error("C-GetDeployNumPerNs 根据ns获取deployment失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeGetDeploymentPerNsErr)
		return
	}

	// 3.resp
	response.RespOK(ctx, "获取每个namespace的deployment数量成功", data)
}
