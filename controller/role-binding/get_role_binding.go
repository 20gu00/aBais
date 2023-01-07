package roleBinding

import (
	k8sClient "github.com/20gu00/aBais/common/k8s-clientset"
	"github.com/20gu00/aBais/common/response"
	param "github.com/20gu00/aBais/model/param/role-binding"
	service "github.com/20gu00/aBais/service/role-binding"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GetRoleBindings(ctx *gin.Context) {
	// 1.参数
	params := new(param.GetRoleBindingInput)
	if err := ctx.Bind(params); err != nil {
		zap.L().Error("C-GetRoleBindings 绑定请求参数失败", zap.Error(err))
		response.RespErr(ctx, response.CodeInvalidParam)
		return
	}

	// 2.service
	client, err := k8sClient.K8s.GetK8sClient(params.Cluster)
	if err != nil {
		zap.L().Error("C-GetRoleBindings 获取k8s的client失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeGetK8sClientErr)
		return
	}

	data, err := service.RoleBinding.GetRoleBindings(client, params.FilterName, params.Namespace, params.Limit, params.Page)
	if err != nil {
		zap.L().Error("C-GetRoleBindings 获取rolebinding列表失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeGetRoleBindingErr)
		return
	}

	// 3.resp
	response.RespOK(ctx, "获取rolebinding列表成功", data)
}
