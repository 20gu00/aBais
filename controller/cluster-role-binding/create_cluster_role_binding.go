package cluster_role_binding

import (
	k8sClient "github.com/20gu00/aBais/common/k8s-clientset"
	"github.com/20gu00/aBais/common/response"
	service "github.com/20gu00/aBais/service/cluster-rolebinding"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 创建role
func CreateClusterRoleBinding(ctx *gin.Context) {
	var (
		clusterRoleBindingCreate = new(service.ClusterRoleBindingCreate)
		err                      error
	)

	// 1.参数
	if err = ctx.ShouldBindJSON(clusterRoleBindingCreate); err != nil {
		zap.L().Error("C-CreateClusterRoleBinding 绑定请求参数失败", zap.Error(err))
		response.RespErr(ctx, response.CodeInvalidParam)
		return
	}
	//指定的cluster的client
	client, err := k8sClient.K8s.GetK8sClient(clusterRoleBindingCreate.Cluster)
	if err != nil {
		zap.L().Error("C-CreateClusterRoleBinding 获取k8s的client失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeGetK8sClientErr)
		return
	}
	if err = service.ClusterRoleBinding.CreateClusterRoleBinding(client, clusterRoleBindingCreate); err != nil {
		zap.L().Error("C-CreateClusterRoleBinding 创建cluster rolebinding失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeCreateclusterRolebindingErr) //  前段可以使用
		return
	}

	// 3.resp
	response.RespOK(ctx, "创建cluster rolebinding成功", nil)
}
