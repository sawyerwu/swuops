package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sawyerwu/swuops/model/response"
)

func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				response.FailWithMessage(fmt.Sprint(err), c)
				return
			}
		}()
		c.Next()
	}
}

