package pvc

import (
	k8sClient "github.com/20gu00/aBais/common/k8s-clientset"
	"github.com/20gu00/aBais/common/response"
	param "github.com/20gu00/aBais/model/param/pvc"
	"github.com/20gu00/aBais/service/pvc"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 获取pvc详情
func GetPvcDetail(ctx *gin.Context) {
	// 1.参数
	params := new(param.GetPvcDetailInput)
	if err := ctx.Bind(params); err != nil {
		zap.L().Error("C-GetPvcDetail 绑定请求参数失败", zap.Error(err))
		response.RespErr(ctx, response.CodeInvalidParam)
		return
	}

	// 2.service
	client, err := k8sClient.K8s.GetK8sClient(params.Cluster)
	if err != nil {
		zap.L().Error("C-GetPvcDetail 获取k8s的client失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeGetK8sClientErr)
		return
	}

	data, err := pvc.Pvc.GetPvcDetail(client, params.PvcName, params.Namespace)
	if err != nil {
		zap.L().Error("C-GetPvcDetail 获取pvc详情失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeGetPvcDetailErr)
		return
	}

	// 3.resp
	response.RespOK(ctx, "获取Pvc详情成功", data)
}
