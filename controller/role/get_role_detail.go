package role

import (
	k8sClient "github.com/20gu00/aBais/common/k8s-clientset"
	"github.com/20gu00/aBais/common/response"
	param "github.com/20gu00/aBais/model/param/role"
	service "github.com/20gu00/aBais/service/role"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GetRoleDetail(ctx *gin.Context) {
	// 1.参数
	params := new(param.GetRoleDetailInput)
	if err := ctx.Bind(params); err != nil {
		zap.L().Error("C-GetRoleDetail 绑定请求参数失败", zap.Error(err))
		response.RespErr(ctx, response.CodeInvalidParam)
		return
	}

	// 2.service
	client, err := k8sClient.K8s.GetK8sClient(params.Cluster)
	if err != nil {
		zap.L().Error("C-GetRoleDetail 获取k8s的client失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeGetK8sClientErr)
		return
	}

	data, err := service.Role.GetRoleDetail(client, params.RoleName, params.Namespace)
	if err != nil {
		zap.L().Error("C-GetRoleDetail 获取role详情失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeGetRoleDetailErr)
		return
	}

	// 3.resp
	response.RespOK(ctx, "获取role详情成功", data)
}
