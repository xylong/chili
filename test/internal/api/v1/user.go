package v1

import (
	"github.com/xylong/chili"
	"github.com/xylong/chili/test/internal/request"
	"net/http"
)

// Register 注册
func Register(ctx *chili.Context) string {
	return "hello"
}

// Login 登录
func Login(ctx *chili.Context, param *request.LoginParam) {
	ctx.JSON(http.StatusOK, param)
}

// Logoff 注销
func Logoff(ctx *chili.Context) {
	ctx.String(http.StatusOK, ctx.Domain())
}

// Me 个人信息
func Me(ctx *chili.Context) {
	ctx.String(http.StatusOK, "个人信息")
}

// Show 用户信息
func Show(ctx *chili.Context) {
	ctx.String(http.StatusOK, "用户信息")
}

// Update 更新个人信息
func Update(ctx *chili.Context) {

}
