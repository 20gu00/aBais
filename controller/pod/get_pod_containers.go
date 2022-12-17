package pod

import (
	k8sClient "github.com/20gu00/aBais/common/k8s-clientset"
	"github.com/20gu00/aBais/common/response"
	param "github.com/20gu00/aBais/model/param/pod"
	"github.com/20gu00/aBais/service/pod"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 获取pod容器
func GetPodContainer(ctx *gin.Context) {
	// 1.参数校验 绑定
	params := new(param.GetPodContainers)
	// GET ctx.Bind
	if err := ctx.Bind(params); err != nil {
		zap.L().Error("C-GetPodContainer 绑定请求参数失败, ", zap.Error(err))
		response.RespErr(ctx, response.CodeInvalidParam)
		return
	}

	// 2.service
	client, err := k8sClient.K8s.GetK8sClient(params.Cluster)
	if err != nil {
		zap.L().Error("C-GetPodContainer 获取k8s的client失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeGetK8sClientErr)
		return
	}

	data, err := pod.Pod.GetPodContainer(client, params.PodName, params.Namespace)
	if err != nil {
		zap.L().Error("C-GetPodContainer 获取pod中的container", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeGetPodContainerErr)
		return
	}

	// 3.resp
	response.RespOK(ctx, "获取Pod容器成功", data)
}
