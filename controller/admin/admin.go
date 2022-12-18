package admin

import (
	"github.com/20gu00/aBais/common/response"
	param "github.com/20gu00/aBais/model/param/admin"
	service "github.com/20gu00/aBais/service/admin"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// admin登录
func Login(ctx *gin.Context) {
	// 结构体指针
	params := new(param.LoginInput)
	if err := ctx.ShouldBindJSON(params); err != nil {
		// msg fkey:fvalue
		// 日志记录后端详细的错误,不讲详细的错误信息返回给前端
		zap.L().Error("Login 绑定请求参数失败, ", zap.Error(err)) // +err.Error()
		response.RespErr(ctx, response.CodeInvalidParam)
		return
	}

	err := service.Login(params.UserName, params.Password)
	if err != nil {
		response.RespInternalErr(ctx, response.CodeServerIntervalErr)
		return
	}

	response.RespOK(ctx, "成功 login", response.CodeSuccess)
}

func Register(ctx *gin.Context) {

}
