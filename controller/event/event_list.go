package event

import (
	"github.com/20gu00/aBais/common/response"
	param "github.com/20gu00/aBais/model/param/event"
	service "github.com/20gu00/aBais/service/event"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 获取event列表 过滤、排序、分页
func GetEventList(ctx *gin.Context) {
	// 1.参数
	params := new(param.ListEventInput)
	if err := ctx.Bind(params); err != nil {
		zap.L().Error("C-GetEventList 绑定请求参数失败", zap.Error(err))
		response.RespErr(ctx, response.CodeInvalidParam)
		return
	}

	// 2.service
	data, err := service.Event.GetList(params.Name, params.Cluster, params.Page, params.Limit)
	if err != nil {
		zap.L().Error("C-GetEventList 获取event列表失败", zap.Error(err))
		response.RespInternalErr(ctx, response.CodeListEventErr)
		return
	}

	// 3.resp
	response.RespOK(ctx, "获取Event列表成功", data)
}
