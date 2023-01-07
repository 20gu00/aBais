package cluster_role_binding

import (
	k8sClient "github.com/20gu00/aBais/common/k8s-clientset"
	"github.com/20gu00/aBais/common/response"
	param "github.com/20gu00/aBais/model/param/cluster-role-binding"
	service "github.com/20gu00/aBais/service/cluster-rolebinding"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func UpdateClusterRoleBinding(ctx *gin.Context) {
	// 1.参数
	params := new(param.UpdateClusterRoleBindingInput)
	// PUT ShouldBindJSON
	if err := ctx.ShouldBindJSON(params); err != nil {
		zap.L().Error("C-UpdateClusterRoleBinding 绑定请求参数失败", zap.Error(err))
		response.RespErr(ctx, response.CodeInvalidParam)
		return
	}

	// 2.service
	client, err := k8sClient.K8s.GetK8sClient(params.Cluster)
	if err != nil {
		zap.L().Error("C-UpdateClusterRoleBinding 获取k8s的client失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeGetK8sClientErr)
		return
	}

	err = service.ClusterRoleBinding.UpdateClusteroleBinding(client, params.Content)
	if err != nil {
		zap.L().Error("C-UpdateClusterRoleBinding 更新cluster rolebinding失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeUpdateclusterRolebindingErr)
		return
	}

	// 3.resp
	response.RespOK(ctx, "更新cluster roleinding成功", nil)
}
