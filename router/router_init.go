package router

import (
	"github.com/20gu00/aBais/common/logger"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	// gin engine
	r := gin.New()
	r.Use(
		logger.GinLogger(),
		logger.GinRecovery(true),
	)

	SetupRouter(r)
	return r
}
