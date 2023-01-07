package pv

import (
	k8sClient "github.com/20gu00/aBais/common/k8s-clientset"
	"github.com/20gu00/aBais/common/response"
	param "github.com/20gu00/aBais/model/param/pv"
	service "github.com/20gu00/aBais/service/pv"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 更新pvc
func UpdatePv(ctx *gin.Context) {
	// 1.参数
	params := new(param.UpdatePvInput)
	// PUT ShouldBindJSON
	if err := ctx.ShouldBindJSON(params); err != nil {
		zap.L().Error("C-UpdatePv 绑定请求参数失败", zap.Error(err))
		response.RespErr(ctx, response.CodeInvalidParam)
		return
	}

	// 2.service
	client, err := k8sClient.K8s.GetK8sClient(params.Cluster)
	if err != nil {
		zap.L().Error("C-UpdatePv 获取k8s的client失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeGetK8sClientErr)
		return
	}

	err = service.Pv.UpdatePv(client, params.Content)
	if err != nil {
		zap.L().Error("C-UpdatePv 更新pv详情失败", zap.Error(err))
		response.RespInternalErr(ctx, response.COdeUpdatePvErr)
		return
	}

	// 3.resp
	response.RespOK(ctx, "更新Pv成功", nil)
}
