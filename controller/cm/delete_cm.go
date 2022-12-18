package cm

import (
	k8sClient "github.com/20gu00/aBais/common/k8s-clientset"
	"github.com/20gu00/aBais/common/response"
	param "github.com/20gu00/aBais/model/param/cm"
	service "github.com/20gu00/aBais/service/cm"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 删除configmap
func DeleteConfigMap(ctx *gin.Context) {
	// 1.参数
	params := new(param.DeleteCmInput)
	// DELETE ShouldBindJSON
	if err := ctx.ShouldBindJSON(params); err != nil {
		zap.L().Error("C-DeleteConfigMap 绑定请求参数失败", zap.Error(err))
		response.RespErr(ctx, response.CodeInvalidParam)
		return
	}

	// 2.service
	client, err := k8sClient.K8s.GetK8sClient(params.Cluster)
	if err != nil {
		zap.L().Error("C-DeleteConfigMap 获取k8s的client失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeGetK8sClientErr)
		return
	}

	err = service.ConfigMap.DeleteConfigMap(client, params.ConfigMapName, params.Namespace)
	if err != nil {
		zap.L().Error("C-DeleteConfigMap 删除cm失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeDeleteCmErr)
		return
	}

	// 3.resp
	response.RespOK(ctx, "删除cm成功", nil)
}
