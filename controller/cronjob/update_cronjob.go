package cronjob

import (
	k8sClient "github.com/20gu00/aBais/common/k8s-clientset"
	"github.com/20gu00/aBais/common/response"
	param "github.com/20gu00/aBais/model/param/cronjob"
	service "github.com/20gu00/aBais/service/cronjob"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func UpdateCronJob(ctx *gin.Context) {
	// 1.参数
	params := new(param.UpdateCronJobInput)
	// PUT ShouldBindJSON
	if err := ctx.ShouldBindJSON(params); err != nil {
		zap.L().Error("C-UpdateCronJob 绑定请求参数失败", zap.Error(err))
		response.RespErr(ctx, response.CodeInvalidParam)
		return
	}

	// 2.service
	client, err := k8sClient.K8s.GetK8sClient(params.Cluster)
	if err != nil {
		zap.L().Error("C-UpdateCronJob 获取k8s的client失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeGetK8sClientErr)
		return
	}

	err = service.CronJob.UpdateCronJob(client, params.Namespace, params.Content)
	if err != nil {
		zap.L().Error("C-UpdateCronJob 更新cronjob失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeUpdateCronJobErr)
		return
	}

	// 3.resp
	response.RespOK(ctx, "更新cronjob成功", nil)
}
