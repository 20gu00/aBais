package namespace

import (
	k8sClient "github.com/20gu00/aBais/common/k8s-clientset"
	"github.com/20gu00/aBais/common/response"
	service "github.com/20gu00/aBais/service/namespace"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreateNs(ctx *gin.Context) {
	var (
		nsCreate = new(service.NsCreate)
		err      error
	)

	// 1.参数
	if err = ctx.ShouldBindJSON(nsCreate); err != nil {
		zap.L().Error("C-CreateNs 绑定请求参数失败", zap.Error(err))
		response.RespErr(ctx, response.CodeInvalidParam)
		return
	}
	//指定的cluster的client
	client, err := k8sClient.K8s.GetK8sClient(nsCreate.Cluster)
	if err != nil {
		zap.L().Error("C-CreateNs 获取k8s的client失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeGetK8sClientErr)
		return
	}
	if err = service.Namespace.CreatePvc(client, nsCreate); err != nil {
		zap.L().Error("C-CreateNs 创建Ns失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeCreateNsErr) //  前段可以使用
		return
	}

	// 3.resp
	response.RespOK(ctx, "创建Ns成功", nil)
}
