package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/xylong/chili/test/internal/request"
	"net/http"
)

// Register 注册
func Register(ctx *gin.Context, param *request.RegisterParam) {
	ctx.JSON(http.StatusOK, param)
}

// Login 登录
func Login(ctx *gin.Context, param *request.LoginParam) {
	ctx.JSON(http.StatusOK, param)
}

// Logoff 注销
func Logoff(ctx *gin.Context) {
	ctx.String(http.StatusOK, "注销")
}
