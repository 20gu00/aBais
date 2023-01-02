package statefulset

import (
	k8sClient "github.com/20gu00/aBais/common/k8s-clientset"
	"github.com/20gu00/aBais/common/response"
	service "github.com/20gu00/aBais/service/statefulset"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreateStatefulset(ctx *gin.Context) {
	var (
		statefulsetCreate = new(service.StatefulsetCreate)
		err               error
	)

	// 1.参数
	if err = ctx.ShouldBindJSON(statefulsetCreate); err != nil {
		zap.L().Error("C-CreateStatefulset 绑定请求参数失败", zap.Error(err))
		response.RespErr(ctx, response.CodeInvalidParam)
		return
	}
	//指定的cluster的client
	client, err := k8sClient.K8s.GetK8sClient(statefulsetCreate.Cluster)
	if err != nil {
		zap.L().Error("C-CreateStatefulset 获取k8s的client失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeGetK8sClientErr)
		return
	}
	if err = service.StatefulSet.CreateDaemonset(client, statefulsetCreate); err != nil {
		zap.L().Error("C-CreateStatefulset 创建Statefulset失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeCreateStatefulsetErr) //  前段可以使用
		return
	}

	// 3.resp
	response.RespOK(ctx, "创建Statefulset成功", nil)
}
