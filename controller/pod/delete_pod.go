package pod

import (
	"net/http"

	k8sClient "github.com/20gu00/aBais/common/k8s-clientset"
	"github.com/20gu00/aBais/common/response"
	param "github.com/20gu00/aBais/model/param/pod"
	"github.com/20gu00/aBais/service/pod"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 删除pod
func DeletePod(ctx *gin.Context) {
	// 1.参数绑定
	params := new(param.DeletePodInput)
	// PUT ShouldBindJSON
	if err := ctx.ShouldBindJSON(params); err != nil {
		zap.L().Error("C-DeletePod 绑定请求参数失败", zap.Error(err))
		response.RespErr(ctx, response.CodeInvalidParam)
		return
	}

	// 2.业务service
	client, err := k8sClient.K8s.GetK8sClient(params.Cluster)
	if err != nil {
		zap.L().Error("C-DeletePod 获取k8s的client失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeGetK8sClientErr)
		return
	}
	err = pod.Pod.DeletePod(client, params.PodName, params.Namespace)
	if err != nil {
		zap.L().Error("C-DeletePod 删除pod失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeDeletePodErr)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}

	// 3.resp
	response.RespOK(ctx, "删除Pod成功", nil)
}
