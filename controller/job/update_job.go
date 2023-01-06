package job

import (
	k8sClient "github.com/20gu00/aBais/common/k8s-clientset"
	"github.com/20gu00/aBais/common/response"
	param "github.com/20gu00/aBais/model/param/job"
	service "github.com/20gu00/aBais/service/job"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func UpdateJob(ctx *gin.Context) {
	// 1.参数
	params := new(param.UpdateJobInput)
	// PUT ShouldBindJSON
	if err := ctx.ShouldBindJSON(params); err != nil {
		zap.L().Error("C-UpdateJob 绑定请求参数失败", zap.Error(err))
		response.RespErr(ctx, response.CodeInvalidParam)
		return
	}

	// 2.service
	client, err := k8sClient.K8s.GetK8sClient(params.Cluster)
	if err != nil {
		zap.L().Error("C-UpdateJob 获取k8s的client失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeGetK8sClientErr)
		return
	}

	err = service.Job.UpdateJob(client, params.Namespace, params.Content)
	if err != nil {
		zap.L().Error("C-UpdateJob 更新job失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeUpdateJobErr)
		return
	}

	// 3.resp
	response.RespOK(ctx, "更新job成功", nil)
}
