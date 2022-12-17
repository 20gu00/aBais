package statefulset

import (
	k8sClient "github.com/20gu00/aBais/common/k8s-clientset"
	"github.com/20gu00/aBais/common/response"
	param "github.com/20gu00/aBais/model/param/statefulset"
	"github.com/20gu00/aBais/service/statefulset"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 删除statefulset
func DeleteStatefulSet(ctx *gin.Context) {
	// 1.参数
	params := new(param.DeleteStatefulsetInput)
	//DELETE请求，绑定参数方法改为ctx.ShouldBindJSON
	if err := ctx.ShouldBindJSON(params); err != nil {
		zap.L().Error("C-DeleteStatefulSet 绑定请求参数失败", zap.Error(err))
		response.RespErr(ctx, response.CodeInvalidParam)
		return
	}

	// 2.service
	client, err := k8sClient.K8s.GetK8sClient(params.Cluster)
	if err != nil {
		zap.L().Error("C-DeleteStatefulSet 获取k8s的client失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeGetK8sClientErr)
		return
	}

	err = statefulset.StatefulSet.DeleteStatefulSet(client, params.StatefulSetName, params.Namespace)
	if err != nil {
		zap.L().Error("C-DeleteStatefulSet 删除statefulset失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeDeleteStatefulsetErr)
		return
	}

	// 3.resp
	response.RespOK(ctx, "删除statefulset成功", nil)
}
