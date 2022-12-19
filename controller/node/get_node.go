package node

import (
	k8sClient "github.com/20gu00/aBais/common/k8s-clientset"
	"github.com/20gu00/aBais/common/response"
	param "github.com/20gu00/aBais/model/param/node"
	service "github.com/20gu00/aBais/service/node"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 获取node列表，支持过滤、排序、分页
func GetNodes(ctx *gin.Context) {
	// 1.参数
	params := new(param.GetNodeInput)
	if err := ctx.Bind(params); err != nil {
		zap.L().Error("C-GetNodes 绑定请求参数失败", zap.Error(err))
		response.RespErr(ctx, response.CodeInvalidParam)
		return
	}

	// 2.service
	client, err := k8sClient.K8s.GetK8sClient(params.Cluster)
	if err != nil {
		zap.L().Error("C-GetNodes 获取k8s的client失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeGetK8sClientErr)
		return
	}

	data, err := service.Node.GetNodes(client, params.FilterName, params.Limit, params.Page)
	if err != nil {
		zap.L().Error("C-GetNodes 获取node列表失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeGetNodeErr)
		return
	}

	// 3.resp
	response.RespOK(ctx, "获取Node列表成功", data)
}
