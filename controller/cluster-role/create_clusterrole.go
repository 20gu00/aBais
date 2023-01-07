package clusterRole

import (
	k8sClient "github.com/20gu00/aBais/common/k8s-clientset"
	"github.com/20gu00/aBais/common/response"
	service "github.com/20gu00/aBais/service/cluster-role"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 创建role
func CreateClusterRole(ctx *gin.Context) {
	var (
		clusterRoleCreate = new(service.ClusterRoleCreate)
		err               error
	)

	// 1.参数
	if err = ctx.ShouldBindJSON(clusterRoleCreate); err != nil {
		zap.L().Error("C-CreateClusterRole 绑定请求参数失败", zap.Error(err))
		response.RespErr(ctx, response.CodeInvalidParam)
		return
	}
	//指定的cluster的client
	client, err := k8sClient.K8s.GetK8sClient(clusterRoleCreate.Cluster)
	if err != nil {
		zap.L().Error("C-CreateClusterRole 获取k8s的client失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeGetK8sClientErr)
		return
	}
	if err = service.ClusterRole.CreateClusterRole(client, clusterRoleCreate); err != nil {
		zap.L().Error("C-CreateClusterRole 创建cluster role失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeCreateClusterRoleErr) //  前段可以使用
		return
	}

	// 3.resp
	response.RespOK(ctx, "创建cluster role成功", nil)
}
