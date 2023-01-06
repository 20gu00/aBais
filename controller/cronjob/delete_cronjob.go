package cronjob

import (
	k8sClient "github.com/20gu00/aBais/common/k8s-clientset"
	"github.com/20gu00/aBais/common/response"
	param "github.com/20gu00/aBais/model/param/cronjob"
	service "github.com/20gu00/aBais/service/cronjob"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func DeleteCronJob(ctx *gin.Context) {
	// 1.参数
	params := new(param.DeleteCronJobInput)
	// DELETE ShouldBindJSON
	if err := ctx.ShouldBindJSON(params); err != nil {
		zap.L().Error("C-DeleteCronJob 绑定请求参数失败", zap.Error(err))
		response.RespErr(ctx, response.CodeInvalidParam)
		return
	}

	// 2.service
	client, err := k8sClient.K8s.GetK8sClient(params.Cluster)
	if err != nil {
		zap.L().Error("C-DeleteCronJob 获取k8s的client失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeGetK8sClientErr)
		return
	}

	err = service.CronJob.DeleteCronJob(client, params.CronJobName, params.Namespace)
	if err != nil {
		zap.L().Error("C-DeleteCronJob 删除CronJob失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeDeleteCronJobErr)
		return
	}

	// 3.resp
	response.RespOK(ctx, "删除CronJob成功", nil)
}
