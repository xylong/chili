package v1

import (
	"github.com/xylong/chili"
	"github.com/xylong/chili/test/internal/request"
	"net/http"
)

// Register 注册
func Register(ctx *chili.Context, param *request.RegisterParam) {
	ctx.JSON(http.StatusOK, param)
}

// Login 登录
func Login(ctx *chili.Context, param *request.LoginParam) {
	ctx.JSON(http.StatusOK, param)
}

// Logoff 注销
func Logoff(ctx *chili.Context) {
	ctx.String(http.StatusOK, ctx.Domain())
}
