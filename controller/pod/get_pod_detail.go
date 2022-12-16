package pod

import (
	"go.uber.org/zap"
	"net/http"

	k8sClient "github.com/20gu00/aBais/common/k8s-clientset"
	"github.com/20gu00/aBais/common/response"
	param "github.com/20gu00/aBais/model/param/pod"

	"github.com/gin-gonic/gin"
)

//获取pod详情
func GetPodDetail(ctx *gin.Context) {
	params := new(param.GetPodDetailInput)
	if err := ctx.Bind(params); err != nil {
		zap.L().Error("GetPodDetail 绑定请求参数失败, ", zap.Error(err))
		response.RespErr(ctx, response.CodeInvalidParam)
		return
	}

	client, err := k8sClient.K8s.GetK8sClient(params.Cluster)
	if err != nil {
		response.RespInternalErr(ctx, response.CodeGetK8sClientErr)
		return
	}

	data, err := service.Pod.GetPodDetail(client, params.PodName, params.Namespace)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "获取Pod详情成功",
		"data": data,
	})
}
