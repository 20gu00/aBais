package roleBinding

import (
	k8sClient "github.com/20gu00/aBais/common/k8s-clientset"
	"github.com/20gu00/aBais/common/response"
	param "github.com/20gu00/aBais/model/param/role-binding"
	service "github.com/20gu00/aBais/service/role-binding"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func DeleteRoleBinding(ctx *gin.Context) {
	// 1.参数
	params := new(param.DeleteRoleBindingInput)
	// DELETE ShouldBindJSON
	if err := ctx.ShouldBindJSON(params); err != nil {
		zap.L().Error("C-DeleteRoleBinding 绑定请求参数失败", zap.Error(err))
		response.RespErr(ctx, response.CodeInvalidParam)
		return
	}

	// 2.service
	client, err := k8sClient.K8s.GetK8sClient(params.Cluster)
	if err != nil {
		zap.L().Error("C-DeleteRoleBinding 获取k8s的client失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeGetK8sClientErr)
		return
	}

	err = service.RoleBinding.DeleteRoleBinding(client, params.RoleBindingName, params.Namespace)
	if err != nil {
		zap.L().Error("C-DeleteRoleBinding 删除RoleBinding失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeDeleteRoleBindingErr)
		return
	}

	// 3.resp
	response.RespOK(ctx, "删除RoleBinding成功", nil)
}
