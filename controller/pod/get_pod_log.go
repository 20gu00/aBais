package pod

import (
	k8sClient "github.com/20gu00/aBais/common/k8s-clientset"
	"github.com/20gu00/aBais/common/response"
	param "github.com/20gu00/aBais/model/param/pod"
	"github.com/20gu00/aBais/service/pod"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 获取pod中容器日志
func GetPodLog(ctx *gin.Context) {
	// 1.参数
	params := new(param.GetPodLog)
	if err := ctx.Bind(params); err != nil {
		zap.L().Error("C-GetPodLog 绑定请求参数失败, ", zap.Error(err))
		response.RespErr(ctx, response.CodeInvalidParam)
		return
	}

	// 2.service
	client, err := k8sClient.K8s.GetK8sClient(params.Cluster)
	if err != nil {
		zap.L().Error("C-GetPodLog 获取k8s的client失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeGetK8sClientErr)
		return
	}

	// string类型,前端在处理
	data, err := pod.Pod.GetPodLog(client, params.ContainerName, params.PodName, params.Namespace)
	if err != nil {
		zap.L().Error("C-GetPodLog 获取pod容器日志失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodePodContainerLogErr)
		return
	}

	// 3.resp
	response.RespOK(ctx, "获取Pod中容器日志成功", data)
}
