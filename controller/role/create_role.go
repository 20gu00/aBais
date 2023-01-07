package role

import (
	k8sClient "github.com/20gu00/aBais/common/k8s-clientset"
	"github.com/20gu00/aBais/common/response"
	service "github.com/20gu00/aBais/service/role"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 创建role
func CreateRole(ctx *gin.Context) {
	var (
		roleCreate = new(service.RoleCreate)
		err        error
	)

	// 1.参数
	if err = ctx.ShouldBindJSON(roleCreate); err != nil {
		zap.L().Error("C-CreateRole 绑定请求参数失败", zap.Error(err))
		response.RespErr(ctx, response.CodeInvalidParam)
		return
	}
	//指定的cluster的client
	client, err := k8sClient.K8s.GetK8sClient(roleCreate.Cluster)
	if err != nil {
		zap.L().Error("C-CreateRole 获取k8s的client失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeGetK8sClientErr)
		return
	}
	if err = service.Role.CreateRole(client, roleCreate); err != nil {
		zap.L().Error("C-CreateRole 创建role失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeCreateRoleErr) //  前段可以使用
		return
	}

	// 3.resp
	response.RespOK(ctx, "创建role成功", nil)
}
