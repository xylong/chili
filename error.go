package chili

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// exception 异常处理
func exception() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.(error).Error()})
			}
		}()

		ctx.Next()
	}
}
