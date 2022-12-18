package daemonset

import (
	k8sClient "github.com/20gu00/aBais/common/k8s-clientset"
	"github.com/20gu00/aBais/common/response"
	param "github.com/20gu00/aBais/model/param/daemonset"
	"github.com/20gu00/aBais/service/daemonset"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GetDaemonSets(ctx *gin.Context) {
	// 1.参数
	params := new(param.GetDaemonsetInput)
	if err := ctx.Bind(params); err != nil {
		zap.L().Error("C-GetDaemonSets 绑定请求参数失败", zap.Error(err))
		response.RespErr(ctx, response.CodeInvalidParam)
		return
	}

	// 2.service
	client, err := k8sClient.K8s.GetK8sClient(params.Cluster)
	if err != nil {
		zap.L().Error("C-GetDaemonSets 获取k8s的client失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeGetK8sClientErr)
		return
	}

	data, err := daemonset.DaemonSet.GetDaemonSets(client, params.FilterName, params.Namespace, params.Limit, params.Page)
	if err != nil {
		zap.L().Error("C-GetDaemonSets 获取daemonset失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeUpdateDaemonsetErr)
		return
	}

	// 3.resp
	response.RespOK(ctx, "获取DaemonSet列表成功", data)
}