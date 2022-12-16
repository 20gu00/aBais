package router

import (
	"github.com/20gu00/aBais/common/logger"
	"github.com/20gu00/aBais/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	// gin engine
	r := gin.New()
	r.Use(
		logger.GinLogger(),
		logger.GinRecovery(true),
		middleware.JWTAuth(),
	)

	SetupRouter(r)
	return r
}
