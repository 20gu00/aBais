package roleBinding

import (
	k8sClient "github.com/20gu00/aBais/common/k8s-clientset"
	"github.com/20gu00/aBais/common/response"
	service "github.com/20gu00/aBais/service/role-binding"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 创建role
func CreateRoleBinding(ctx *gin.Context) {
	var (
		roleBindingCreate = new(service.RoleBindingCreate)
		err               error
	)

	// 1.参数
	if err = ctx.ShouldBindJSON(roleBindingCreate); err != nil {
		zap.L().Error("C-CreateRoleBinding 绑定请求参数失败", zap.Error(err))
		response.RespErr(ctx, response.CodeInvalidParam)
		return
	}
	//指定的cluster的client
	client, err := k8sClient.K8s.GetK8sClient(roleBindingCreate.Cluster)
	if err != nil {
		zap.L().Error("C-CreateRoleBinding 获取k8s的client失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeGetK8sClientErr)
		return
	}
	if err = service.RoleBinding.CreateRoleBinding(client, roleBindingCreate); err != nil {
		zap.L().Error("C-CreateRoleBinding 创建rolebinding失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeCreateRoleBindingErr) //  前段可以使用
		return
	}

	// 3.resp
	response.RespOK(ctx, "创建rolebinding成功", nil)
}
