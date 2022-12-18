package pvc

import (
	k8sClient "github.com/20gu00/aBais/common/k8s-clientset"
	"github.com/20gu00/aBais/common/response"
	param "github.com/20gu00/aBais/model/param/pvc"
	"github.com/20gu00/aBais/service/pvc"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 删除pvc
func DeletePvc(ctx *gin.Context) {
	// 1.参数
	params := new(param.DeletePvcInput)
	//DELETE请求，绑定参数方法改为ctx.ShouldBindJSON
	if err := ctx.ShouldBindJSON(params); err != nil {
		zap.L().Error("C-DeletePvc 绑定请求参数失败", zap.Error(err))
		response.RespErr(ctx, response.CodeInvalidParam)
		return
	}

	// 2.service
	client, err := k8sClient.K8s.GetK8sClient(params.Cluster)
	if err != nil {
		zap.L().Error("C-DeletePvc 获取k8s的client失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeGetK8sClientErr)
		return
	}

	err = pvc.Pvc.DeletePvc(client, params.PvcName, params.Namespace)
	if err != nil {
		zap.L().Error("C-DeletePvc 删除pvc失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeGetPvcErr)
		return
	}

	// 3.resp
	response.RespOK(ctx, "删除Pvc成功", nil)
}
