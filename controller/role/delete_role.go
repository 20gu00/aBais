package role

import (
	k8sClient "github.com/20gu00/aBais/common/k8s-clientset"
	"github.com/20gu00/aBais/common/response"
	param "github.com/20gu00/aBais/model/param/role"
	service "github.com/20gu00/aBais/service/role"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func DeleteRole(ctx *gin.Context) {
	// 1.参数
	params := new(param.DeleteRoleInput)
	// DELETE ShouldBindJSON
	if err := ctx.ShouldBindJSON(params); err != nil {
		zap.L().Error("C-DeleteRole 绑定请求参数失败", zap.Error(err))
		response.RespErr(ctx, response.CodeInvalidParam)
		return
	}

	// 2.service
	client, err := k8sClient.K8s.GetK8sClient(params.Cluster)
	if err != nil {
		zap.L().Error("C-DeleteRole 获取k8s的client失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeGetK8sClientErr)
		return
	}

	err = service.Role.DeleteRole(client, params.RoleName, params.Namespace)
	if err != nil {
		zap.L().Error("C-DeleteRole 删除Role失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeDeleteRoleErr)
		return
	}

	// 3.resp
	response.RespOK(ctx, "删除Role成功", nil)
}
