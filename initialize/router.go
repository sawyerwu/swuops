package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/sawyerwu/swuops/middleware"
)

func InitRouters() *gin.Engine {
	r := gin.New()
	r.Use(middleware.LoggerToFile(), middleware.RecoveryMiddleware())

	return r
}
