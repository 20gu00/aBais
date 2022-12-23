package pod

import (
	k8sClient "github.com/20gu00/aBais/common/k8s-clientset"
	"github.com/20gu00/aBais/common/response"
	service "github.com/20gu00/aBais/service/pod"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreatePod(ctx *gin.Context) {
	var (
		podCreateParam = new(service.PodCreateParam)
		err            error
	)

	if err = ctx.ShouldBindJSON(podCreateParam); err != nil {
		zap.L().Error("C-CreatePod 绑定请求参数失败", zap.Error(err))
		response.RespErr(ctx, response.CodeInvalidParam)
		return
	}
	client, err := k8sClient.K8s.GetK8sClient(podCreateParam.Cluster)
	if err != nil {
		zap.L().Error("C-CreatePod 获取k8s的client失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeGetK8sClientErr)
		return
	}
	if err = service.Pod.CreatePod(client, podCreateParam); err != nil {
		zap.L().Error("C-CreatePod 创建pod失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeCreatePodErr)
		return
	}

	// 3.resp
	response.RespOK(ctx, "创建pod成功", nil)
}
