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

// 获取每个namespace的pod数量
func GetPodNumPerNs(ctx *gin.Context) {
	// 1.参数
	params := new(param.GetPodNumPerNamespace)
	//GET请求，绑定参数方法改为ctx.Bind
	if err := ctx.Bind(params); err != nil {
		zap.L().Error("C-GetPodNumPerNp 绑定请求参数失败, ", zap.Error(err))
		response.RespErr(ctx, response.CodeInvalidParam)
		return
	}

	// 2.service
	client, err := k8sClient.K8s.GetK8sClient(params.Cluster)
	if err != nil {
		zap.L().Error("C-GetPodNumPerNp 获取k8s的client失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeGetK8sClientErr)
		return
	}

	data, err := pod.Pod.GetPodNumPerNs(client)
	if err != nil {
		zap.L().Error("C-GetPodNumPerNp 根据namespace获取pod数目失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeGetNumByNsErr)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}

	// 3.resp
	response.RespOK(ctx, "获取每个namespace的pod数量成功", data)
}
