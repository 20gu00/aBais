package job

import (
	k8sClient "github.com/20gu00/aBais/common/k8s-clientset"
	"github.com/20gu00/aBais/common/response"
	service "github.com/20gu00/aBais/service/job"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 创建job
func CreateJob(ctx *gin.Context) {
	var (
		jobCreate = new(service.JobCreate)
		err       error
	)

	// 1.参数
	if err = ctx.ShouldBindJSON(jobCreate); err != nil {
		zap.L().Error("C-CreateJob 绑定请求参数失败", zap.Error(err))
		response.RespErr(ctx, response.CodeInvalidParam)
		return
	}
	//指定的cluster的client
	client, err := k8sClient.K8s.GetK8sClient(jobCreate.Cluster)
	if err != nil {
		zap.L().Error("C-CreateJob 获取k8s的client失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeGetK8sClientErr)
		return
	}
	if err = service.Job.CreateJob(client, jobCreate); err != nil {
		zap.L().Error("C-CreateJob 创建job失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeCreateJobErr) //  前段可以使用
		return
	}

	// 3.resp
	response.RespOK(ctx, "创建Job成功", nil)
}
