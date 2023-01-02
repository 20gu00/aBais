package daemonset

import (
	k8sClient "github.com/20gu00/aBais/common/k8s-clientset"
	"github.com/20gu00/aBais/common/response"
	service "github.com/20gu00/aBais/service/daemonset"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreateDaemonset(ctx *gin.Context) {
	var (
		daemonsetCreate = new(service.DaemonsetCreate)
		err             error
	)

	// 1.参数
	if err = ctx.ShouldBindJSON(daemonsetCreate); err != nil {
		zap.L().Error("C-CreateDaemonset 绑定请求参数失败", zap.Error(err))
		response.RespErr(ctx, response.CodeInvalidParam)
		return
	}
	//指定的cluster的client
	client, err := k8sClient.K8s.GetK8sClient(daemonsetCreate.Cluster)
	if err != nil {
		zap.L().Error("C-CreateDaemonset 获取k8s的client失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeGetK8sClientErr)
		return
	}
	if err = service.DaemonSet.CreateDaemonset(client, daemonsetCreate); err != nil {
		zap.L().Error("C-CreateDaemonset 创建daemonset失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeCreateDaemonsetErr) //  前段可以使用
		return
	}

	// 3.resp
	response.RespOK(ctx, "创建Daemonset成功", nil)
}
