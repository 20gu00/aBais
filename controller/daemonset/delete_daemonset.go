package daemonset

import (
	k8sClient "github.com/20gu00/aBais/common/k8s-clientset"
	"github.com/20gu00/aBais/common/response"
	param "github.com/20gu00/aBais/model/param/daemonset"
	"github.com/20gu00/aBais/service/daemonset"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 删除daemonset
func DeleteDaemonSet(ctx *gin.Context) {
	// 1.参数
	params := new(param.DeleteDaemonsetInput)
	//DELETE请求，绑定参数方法改为ctx.ShouldBindJSON
	if err := ctx.ShouldBindJSON(params); err != nil {
		zap.L().Error("C-DeleteDaemonSet 绑定请求参数失败", zap.Error(err))
		response.RespErr(ctx, response.CodeInvalidParam)
		return
	}

	// 2.service
	client, err := k8sClient.K8s.GetK8sClient(params.Cluster)
	if err != nil {
		zap.L().Error("C-DeleteDaemonSet 获取k8s的client失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeGetK8sClientErr)
		return
	}

	err = daemonset.DaemonSet.DeleteDaemonSet(client, params.DaemonSetName, params.Namespace)
	if err != nil {
		zap.L().Error("C-DeleteDaemonSet 删除daemonset失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeDeleteDaemonsetErr)
		return
	}

	// 3.resp
	response.RespOK(ctx, "删除DaemonSet成功", nil)
}
