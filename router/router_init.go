package router

import (
	"github.com/20gu00/aBais/common/logger"
	"github.com/20gu00/aBais/common/response"
	"github.com/20gu00/aBais/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	// gin engine
	r := gin.New()
	// ping
	r.GET("/ping", func(ctx *gin.Context) {
		response.RespOK(ctx, "ping测试成功", nil)
	})

	r.Use(
		logger.GinLogger(),
		logger.GinRecovery(true),
		middleware.JWTAuth(),
		middleware.Cors(),
	)

	SetupRouter(r)
	return r
}
