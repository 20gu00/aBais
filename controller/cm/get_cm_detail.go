package cm

import (
	k8sClient "github.com/20gu00/aBais/common/k8s-clientset"
	"github.com/20gu00/aBais/common/response"
	param "github.com/20gu00/aBais/model/param/cm"
	"github.com/20gu00/aBais/service/cm"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 获取configmap详情
func GetConfigMapDetail(ctx *gin.Context) {
	// 1.参数
	params := new(param.GetCmDetailInput)
	if err := ctx.Bind(params); err != nil {
		zap.L().Error("C-GetConfigMapDetail 绑定请求参数失败", zap.Error(err))
		response.RespErr(ctx, response.CodeInvalidParam)
		return
	}

	// 2.service
	client, err := k8sClient.K8s.GetK8sClient(params.Cluster)
	if err != nil {
		zap.L().Error("C-GetConfigMapDetail 获取k8s的client失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeGetK8sClientErr)
		return
	}

	data, err := cm.ConfigMap.GetConfigMapDetail(client, params.ConfigMapName, params.Namespace)
	if err != nil {
		zap.L().Error("C-GetConfigMapDetail 获取cm详情失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeGetCmDetailErr)
		return
	}

	// 3.resp
	response.RespOK(ctx, "获取cm详情成功", data)
}
