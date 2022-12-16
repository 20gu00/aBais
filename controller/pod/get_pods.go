package pod

import (
	k8sClient "github.com/20gu00/aBais/common/k8s-clientset"
	"github.com/20gu00/aBais/common/response"
	param "github.com/20gu00/aBais/model/param/pod"
	"github.com/20gu00/aBais/service/pod"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GetPods(ctx *gin.Context) {
	// 1.请求参数绑定
	params := new(param.GetPodsInput)
	// form格式 ctx.Bind
	if err := ctx.Bind(params); err != nil {
		zap.L().Error("GetPods 绑定请求参数失败, ", zap.Error(err))
		response.RespErr(ctx, response.CodeInvalidParam)
		return
	}

	// 2.service
	// 获取k8s的client
	client, err := k8sClient.K8s.GetK8sClient(params.Cluster)
	if err != nil {
		response.RespInternalErr(ctx, response.CodeGetK8sClientErr)
		return
	}
	// 获取pods
	data, err := pod.Pod.GetPods(client, params.FilterName, params.Namespace, params.Limit, params.Page)
	if err != nil {
		response.RespInternalErr(ctx, response.CodeGetPodListErr)
		return
	}

	// 3.resp
	response.RespOK(ctx, "获取Pod列表成功", data)
}
