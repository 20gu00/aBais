package serviceAccount

import (
	k8sClient "github.com/20gu00/aBais/common/k8s-clientset"
	"github.com/20gu00/aBais/common/response"
	param "github.com/20gu00/aBais/model/param/sa"
	service "github.com/20gu00/aBais/service/service-account"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GetSas(ctx *gin.Context) {
	// 1.参数
	params := new(param.GetSaInput)
	if err := ctx.Bind(params); err != nil {
		zap.L().Error("C-GetSas 绑定请求参数失败", zap.Error(err))
		response.RespErr(ctx, response.CodeInvalidParam)
		return
	}

	// 2.service
	client, err := k8sClient.K8s.GetK8sClient(params.Cluster)
	if err != nil {
		zap.L().Error("C-GetSas 获取k8s的client失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeGetK8sClientErr)
		return
	}

	data, err := service.Sa.GetSas(client, params.FilterName, params.Namespace, params.Limit, params.Page)
	if err != nil {
		zap.L().Error("C-GetSas 获取sa列表失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeGetserviceAccountErr)
		return
	}

	// 3.resp
	response.RespOK(ctx, "获取sa列表成功", data)
}
