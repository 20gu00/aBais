package pod

import (
	k8sClient "github.com/20gu00/aBais/common/k8s-clientset"
	"github.com/20gu00/aBais/common/response"
	param "github.com/20gu00/aBais/model/param/pod"
	"github.com/20gu00/aBais/service/pod"
	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

//获取pod详情
func GetPodDetail(ctx *gin.Context) {
	// 1.参数绑定
	params := new(param.GetPodDetailInput)
	if err := ctx.Bind(params); err != nil {
		zap.L().Error("C-GetPodDetail 绑定请求参数失败", zap.Error(err))
		response.RespErr(ctx, response.CodeInvalidParam)
		return
	}

	// 2.业务逻辑service
	client, err := k8sClient.K8s.GetK8sClient(params.Cluster)
	if err != nil {
		zap.L().Error("C-GetPodDetail 获取k8s的client失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeGetK8sClientErr)
		return
	}

	data, err := pod.Pod.GetPodDetail(client, params.PodName, params.Namespace)
	if err != nil {
		zap.L().Error("C-GetPodDetail 获取pod详情信息失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeGetPodDetailErr)
		return
	}

	// 3.resp
	response.RespOK(ctx, "获取Pod详情成功", data)
}
