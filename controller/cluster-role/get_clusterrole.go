package clusterRole

import (
	k8sClient "github.com/20gu00/aBais/common/k8s-clientset"
	"github.com/20gu00/aBais/common/response"
	param "github.com/20gu00/aBais/model/param/cluster-role"
	service "github.com/20gu00/aBais/service/cluster-role"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GetClusterRoles(ctx *gin.Context) {
	// 1.参数
	params := new(param.GetClusterRoleInput)
	if err := ctx.Bind(params); err != nil {
		zap.L().Error("C-GetClusterRoles 绑定请求参数失败", zap.Error(err))
		response.RespErr(ctx, response.CodeInvalidParam)
		return
	}

	// 2.service
	client, err := k8sClient.K8s.GetK8sClient(params.Cluster)
	if err != nil {
		zap.L().Error("C-GetClusterRoles 获取k8s的client失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeGetK8sClientErr)
		return
	}

	data, err := service.ClusterRole.GetClusterRoles(client, params.FilterName, params.Limit, params.Page)
	if err != nil {
		zap.L().Error("C-GetClusterRoles 获取cluster role列表失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeGetClusterRoleErr)
		return
	}

	// 3.resp
	response.RespOK(ctx, "获取cluster role列表成功", data)
}
