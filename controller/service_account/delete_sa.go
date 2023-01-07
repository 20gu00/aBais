package serviceAccount

import (
	k8sClient "github.com/20gu00/aBais/common/k8s-clientset"
	"github.com/20gu00/aBais/common/response"
	param "github.com/20gu00/aBais/model/param/sa"
	service "github.com/20gu00/aBais/service/service-account"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func DeleteSa(ctx *gin.Context) {
	// 1.参数
	params := new(param.DeleteSaInput)
	// DELETE ShouldBindJSON
	if err := ctx.ShouldBindJSON(params); err != nil {
		zap.L().Error("C-DeleteSa 绑定请求参数失败", zap.Error(err))
		response.RespErr(ctx, response.CodeInvalidParam)
		return
	}

	// 2.service
	client, err := k8sClient.K8s.GetK8sClient(params.Cluster)
	if err != nil {
		zap.L().Error("C-DeleteSa 获取k8s的client失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeGetK8sClientErr)
		return
	}

	err = service.Sa.DeleteSa(client, params.SaName, params.Namespace)
	if err != nil {
		zap.L().Error("C-DeleteSa 删除sa失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeDeleteserviceAccountErr)
		return
	}

	// 3.resp
	response.RespOK(ctx, "删除sa成功", nil)
}
