package cm

import (
	k8sClient "github.com/20gu00/aBais/common/k8s-clientset"
	"github.com/20gu00/aBais/common/response"
	service "github.com/20gu00/aBais/service/cm"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreateCm(ctx *gin.Context) {
	var (
		cmCreate = new(service.ConfigmapCreate)
		err      error
	)

	// 1.参数
	if err = ctx.ShouldBindJSON(cmCreate); err != nil {
		zap.L().Error("C-CreateCm 绑定请求参数失败", zap.Error(err))
		response.RespErr(ctx, response.CodeInvalidParam)
		return
	}
	//指定的cluster的client
	client, err := k8sClient.K8s.GetK8sClient(cmCreate.Cluster)
	if err != nil {
		zap.L().Error("C-CreateCm 获取k8s的client失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeGetK8sClientErr)
		return
	}
	if err = service.ConfigMap.CreateCm(client, cmCreate); err != nil {
		zap.L().Error("C-CreateConfigmap 创建Configmap失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeCreateCmErr) //  前段可以使用
		return
	}

	// 3.resp
	response.RespOK(ctx, "创建Configmap成功", nil)
}
