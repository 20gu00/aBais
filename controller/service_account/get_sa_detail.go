package serviceAccount

import (
	k8sClient "github.com/20gu00/aBais/common/k8s-clientset"
	"github.com/20gu00/aBais/common/response"
	param "github.com/20gu00/aBais/model/param/sa"
	service "github.com/20gu00/aBais/service/service-account"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GetSaDetail(ctx *gin.Context) {
	// 1.参数
	params := new(param.GetSaDetailInput)
	if err := ctx.Bind(params); err != nil {
		zap.L().Error("C-GetSaDetail 绑定请求参数失败", zap.Error(err))
		response.RespErr(ctx, response.CodeInvalidParam)
		return
	}

	// 2.service
	client, err := k8sClient.K8s.GetK8sClient(params.Cluster)
	if err != nil {
		zap.L().Error("C-GetSaDetail 获取k8s的client失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeGetK8sClientErr)
		return
	}

	data, err := service.Sa.GetSaDetail(client, params.SaName, params.Namespace)
	if err != nil {
		zap.L().Error("C-GetSaDetail 获取sa详情失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeGetSecretDetailErr)
		return
	}

	// 3.resp
	response.RespOK(ctx, "获取sa详情成功", data)
}
