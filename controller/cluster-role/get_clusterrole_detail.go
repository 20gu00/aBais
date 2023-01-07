package clusterRole

import (
	k8sClient "github.com/20gu00/aBais/common/k8s-clientset"
	"github.com/20gu00/aBais/common/response"
	param "github.com/20gu00/aBais/model/param/cluster-role"
	service "github.com/20gu00/aBais/service/cluster-role"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GetClusterRoleDetail(ctx *gin.Context) {
	// 1.参数
	params := new(param.GetClusterRoleDetailInput)
	if err := ctx.Bind(params); err != nil {
		zap.L().Error("C-GetClusterRoleDetail 绑定请求参数失败", zap.Error(err))
		response.RespErr(ctx, response.CodeInvalidParam)
		return
	}

	// 2.service
	client, err := k8sClient.K8s.GetK8sClient(params.Cluster)
	if err != nil {
		zap.L().Error("C-GetClusterRoleDetail 获取k8s的client失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeGetK8sClientErr)
		return
	}

	data, err := service.ClusterRole.GetClusterRoleDetail(client, params.ClusterRoleName)
	if err != nil {
		zap.L().Error("C-GetClusterRoleDetail 获取cluster role详情失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeGetClusterRoleDetailErr)
		return
	}

	// 3.resp
	response.RespOK(ctx, "获取cluster role详情成功", data)
}
