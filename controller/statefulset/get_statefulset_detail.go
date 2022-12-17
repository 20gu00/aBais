package statefulset

import (
	k8sClient "github.com/20gu00/aBais/common/k8s-clientset"
	"github.com/20gu00/aBais/common/response"
	param "github.com/20gu00/aBais/model/param/statefulset"
	"github.com/20gu00/aBais/service/statefulset"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 获取statefulset详情
func GetStatefulSetDetail(ctx *gin.Context) {
	// 1.参数
	params := new(param.GetStatefulsetDetailInput)
	if err := ctx.Bind(params); err != nil {
		zap.L().Error("C-GetStatefulSetDetail 绑定请求参数失败", zap.Error(err))
		response.RespErr(ctx, response.CodeInvalidParam)
		return
	}

	// 2.service
	client, err := k8sClient.K8s.GetK8sClient(params.Cluster)
	if err != nil {
		zap.L().Error("C-GetStatefulSetDetail 获取k8s的client失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeGetK8sClientErr)
		return
	}
	data, err := statefulset.StatefulSet.GetStatefulSetDetail(client, params.StatefulSetName, params.Namespace)
	if err != nil {
		zap.L().Error("C-GetStatefulSetDetail 获取statefulset列表失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeGetStatefulsetDetailErr)
		return
	}

	// 3.resp
	response.RespOK(ctx, "获取StatefulSet详情成功", data)
}
