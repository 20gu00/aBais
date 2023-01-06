package cronjob

import (
	k8sClient "github.com/20gu00/aBais/common/k8s-clientset"
	"github.com/20gu00/aBais/common/response"
	service "github.com/20gu00/aBais/service/cronjob"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 创建job
func CreateCronJob(ctx *gin.Context) {
	var (
		cronJobCreate = new(service.CronJobCreate)
		err           error
	)

	// 1.参数
	if err = ctx.ShouldBindJSON(cronJobCreate); err != nil {
		zap.L().Error("C-CreateCronJob 绑定请求参数失败", zap.Error(err))
		response.RespErr(ctx, response.CodeInvalidParam)
		return
	}
	//指定的cluster的client
	client, err := k8sClient.K8s.GetK8sClient(cronJobCreate.Cluster)
	if err != nil {
		zap.L().Error("C-CreateCronJob 获取k8s的client失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeGetK8sClientErr)
		return
	}
	if err = service.CronJob.CreateCronJob(client, cronJobCreate); err != nil {
		zap.L().Error("C-CreateCronJob 创建cronJob失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeCreateCronJobErr) //  前段可以使用
		return
	}

	// 3.resp
	response.RespOK(ctx, "创建cronJob成功", nil)
}
