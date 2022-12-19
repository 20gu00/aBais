package node

import (
	k8sClient "github.com/20gu00/aBais/common/k8s-clientset"
	"github.com/20gu00/aBais/common/response"
	param "github.com/20gu00/aBais/model/param/node"
	service "github.com/20gu00/aBais/service/node"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 获取node详情
func GetNodeDetail(ctx *gin.Context) {
	// 1.参数
	params := new(param.GetNodeDetailInput)
	if err := ctx.Bind(params); err != nil {
		zap.L().Error("C-GetNodeDetail 绑定请求参数失败", zap.Error(err))
		response.RespErr(ctx, response.CodeInvalidParam)
		return
	}

	// 2.service
	client, err := k8sClient.K8s.GetK8sClient(params.Cluster)
	if err != nil {
		zap.L().Error("C-GetNodeDetail 获取k8s的client失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeGetK8sClientErr)
		return
	}

	data, err := service.Node.GetNodeDetail(client, params.NodeName)
	if err != nil {
		zap.L().Error("C-GGetNodeDetailetNodes 获取node详情失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeGetNodeDetailErr)
		return
	}

	// 3.resp
	response.RespOK(ctx, "获取Node详情成功", data)
}
