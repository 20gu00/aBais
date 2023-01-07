package roleBinding

import (
	k8sClient "github.com/20gu00/aBais/common/k8s-clientset"
	"github.com/20gu00/aBais/common/response"
	param "github.com/20gu00/aBais/model/param/role-binding"
	service "github.com/20gu00/aBais/service/role-binding"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func UpdateRoleBinding(ctx *gin.Context) {
	// 1.参数
	params := new(param.UpdateRoleBindingInput)
	// PUT ShouldBindJSON
	if err := ctx.ShouldBindJSON(params); err != nil {
		zap.L().Error("C-UpdateRoleBinding 绑定请求参数失败", zap.Error(err))
		response.RespErr(ctx, response.CodeInvalidParam)
		return
	}

	// 2.service
	client, err := k8sClient.K8s.GetK8sClient(params.Cluster)
	if err != nil {
		zap.L().Error("C-UpdateRoleBinding 获取k8s的client失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeGetK8sClientErr)
		return
	}

	err = service.RoleBinding.UpdateRoleBinding(client, params.Namespace, params.Content)
	if err != nil {
		zap.L().Error("C-UpdateRoleBinding 更新rolebinding失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeUpdateRoleBindingErr)
		return
	}

	// 3.resp
	response.RespOK(ctx, "更新roleinding成功", nil)
}
