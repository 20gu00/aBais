package pod

import (
	k8sClient "github.com/20gu00/aBais/common/k8s-clientset"
	"github.com/20gu00/aBais/common/response"
	param "github.com/20gu00/aBais/model/param/pod"
	"github.com/20gu00/aBais/service/pod"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 更新pod
func UpdatePod(ctx *gin.Context) {
	// 1.参数校验 绑定
	params := new(param.UpdatePodInput)
	//PUT请求，绑定参数方法改为ctx.ShouldBindJSON
	if err := ctx.ShouldBindJSON(params); err != nil {
		zap.L().Error("C-UpdatePod 绑定请求参数失败", zap.Error(err))
		response.RespErr(ctx, response.CodeInvalidParam)
		return
	}
	// 2.service
	client, err := k8sClient.K8s.GetK8sClient(params.Cluster)
	if err != nil {
		zap.L().Error("C-UpdatePod 获取k8s的client失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeGetK8sClientErr)
		return
	}
	err = pod.Pod.UpdatePod(client, params.PodName, params.Namespace, params.Content)
	if err != nil {
		zap.L().Error("C-UpdatePod 更新pod失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeUpdatePodErr)
		return
	}

	// 3.resp
	response.RespOK(ctx, "更新Pod成功", nil)
}
