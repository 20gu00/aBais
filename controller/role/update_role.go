package role

import (
	k8sClient "github.com/20gu00/aBais/common/k8s-clientset"
	"github.com/20gu00/aBais/common/response"
	param "github.com/20gu00/aBais/model/param/role"
	service "github.com/20gu00/aBais/service/role"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func UpdateRole(ctx *gin.Context) {
	// 1.参数
	params := new(param.UpdateRoleInput)
	// PUT ShouldBindJSON
	if err := ctx.ShouldBindJSON(params); err != nil {
		zap.L().Error("C-UpdateRole 绑定请求参数失败", zap.Error(err))
		response.RespErr(ctx, response.CodeInvalidParam)
		return
	}

	// 2.service
	client, err := k8sClient.K8s.GetK8sClient(params.Cluster)
	if err != nil {
		zap.L().Error("C-UpdateRole 获取k8s的client失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeGetK8sClientErr)
		return
	}

	err = service.Role.UpdateRole(client, params.Namespace, params.Content)
	if err != nil {
		zap.L().Error("C-UpdateRole 更新role失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeUpdateRoleErr)
		return
	}

	// 3.resp
	response.RespOK(ctx, "更新role成功", nil)
}
