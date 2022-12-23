package allResources

import (
	k8sClient "github.com/20gu00/aBais/common/k8s-clientset"
	"github.com/20gu00/aBais/common/response"
	param "github.com/20gu00/aBais/model/param/all-resource"
	service "github.com/20gu00/aBais/service/all-resource"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GetAllResourceNum(ctx *gin.Context) {
	// 1.参数
	params := new(param.GetAllResourceNum)
	if err := ctx.Bind(params); err != nil {
		zap.L().Error("C-GetAllResourceNum 绑定请求参数失败", zap.Error(err))
		response.RespErr(ctx, response.CodeInvalidParam)
		return
	}

	// 2.service
	client, err := k8sClient.K8s.GetK8sClient(params.Cluster)
	if err != nil {
		zap.L().Error("C-GetAllResourceNum 获取k8s的client失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeGetK8sClientErr)
		return
	}

	data, errs := service.AllRes.GetAllResourceNum(client)
	if len(errs) > 0 {
		zap.L().Error("C-GetAllResourceNum 获取所有资源的数量失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeGetAllResourceNumErr)
		return
	}

	// 3.resp
	response.RespOK(ctx, "获取资源数量成功", data)
}
