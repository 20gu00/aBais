package pv

import (
	k8sClient "github.com/20gu00/aBais/common/k8s-clientset"
	"github.com/20gu00/aBais/common/response"
	service "github.com/20gu00/aBais/service/pv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreatePv(ctx *gin.Context) {
	var (
		pvCreate = new(service.PvCreate)
		err      error
	)

	// 1.参数
	if err = ctx.ShouldBindJSON(pvCreate); err != nil {
		zap.L().Error("C-CreatePv 绑定请求参数失败", zap.Error(err))
		response.RespErr(ctx, response.CodeInvalidParam)
		return
	}
	//指定的cluster的client
	client, err := k8sClient.K8s.GetK8sClient(pvCreate.Cluster)
	if err != nil {
		zap.L().Error("C-CreatePv 获取k8s的client失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeGetK8sClientErr)
		return
	}
	if err = service.Pv.CreatePvc(client, pvCreate); err != nil {
		zap.L().Error("C-CreatePv 创建Pvc失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeCreatePvErr) //  前段可以使用
		return
	}

	// 3.resp
	response.RespOK(ctx, "创建Pv成功", nil)
}
