package request

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/xylong/chili"
)

// RegisterParam 注册参数
type RegisterParam struct {
	Phone    string `json:"phone" form:"phone" binding:"required"`
	Password string `json:"password" form:"password" binding:"required,min=6,max=18"`
}

// LoginParam 登录参数
type LoginParam struct {
	Name     string `json:"name" form:"name" binding:"required,alphanum"`
	Password string `json:"password" form:"password" binding:"required,min=6,max=18"`
}

func (p *LoginParam) Bind(ctx *chili.Context) error {
	return ctx.ShouldBindWith(p, binding.Form)
}
