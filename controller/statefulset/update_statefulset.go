package statefulset

import (
	k8sClient "github.com/20gu00/aBais/common/k8s-clientset"
	"github.com/20gu00/aBais/common/response"
	param "github.com/20gu00/aBais/model/param/statefulset"
	"github.com/20gu00/aBais/service/statefulset"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 更新statefulSet
func UpdateStatefulSet(ctx *gin.Context) {
	// 1.参数
	params := new(param.UpdateStatefulsetInput)
	// PUTShouldBindJSON
	if err := ctx.ShouldBindJSON(params); err != nil {
		zap.L().Error("C-UpdateStatefulSet 绑定请求参数失败", zap.Error(err))
		response.RespErr(ctx, response.CodeInvalidParam)
		return
	}

	// 2.service
	client, err := k8sClient.K8s.GetK8sClient(params.Cluster)
	if err != nil {
		zap.L().Error("C-UpdateStatefulSet 获取k8s的client失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeGetK8sClientErr)
		return
	}

	err = statefulset.StatefulSet.UpdateStatefulSet(client, params.Namespace, params.Content)
	if err != nil {
		zap.L().Error("C-UpdateStatefulSet 更新statefulset失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeUpdateStatefulsetErr)
		return
	}

	// 3.resp
	response.RespOK(ctx, "更新StatefulSet成功", nil)
}
