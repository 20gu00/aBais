package namespace

import (
	k8sClient "github.com/20gu00/aBais/common/k8s-clientset"
	"github.com/20gu00/aBais/common/response"
	param "github.com/20gu00/aBais/model/param/ns"
	service "github.com/20gu00/aBais/service/namespace"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 获取namespace列表 过滤、排序、分页
func GetNamespaces(ctx *gin.Context) {
	// 1.参数
	params := new(param.GetNsInput)
	if err := ctx.Bind(params); err != nil {
		zap.L().Error("C-GetNamespaces 绑定请求参数失败", zap.Error(err))
		response.RespErr(ctx, response.CodeInvalidParam)
		return
	}

	// 2.service
	client, err := k8sClient.K8s.GetK8sClient(params.Cluster)
	if err != nil {
		zap.L().Error("C-GetNamespaces 获取k8s的client失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeGetK8sClientErr)
		return
	}

	data, err := service.Namespace.GetNamespaces(client, params.FilterName, params.Limit, params.Page)
	if err != nil {
		zap.L().Error("C-GetNamespaces 获取ns失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeGetNsErr)
		return
	}

	// 3.resp
	response.RespOK(ctx, "获取ns成功", data)
}
