package router

import (
	"github.com/xylong/chili"
	"github.com/xylong/chili/test/internal/api/v1"
	"github.com/xylong/chili/test/internal/dto"
)

var (
	userController *v1.UserController
)

func init() {
	userController = v1.NewUserController()
}

// V1 version 1 routes
func V1(v1 *chili.Group) {
	v1.GET("users/:id", userController.ShowAccount)
	v1.POST("register", userController.Register)
	v1.POST("login", chili.NewBind[dto.LoginParam]().Try(userController.Login).
		Cache(func(ctx *chili.Context, param *dto.LoginParam, err error) {
			ctx.AbortWithStatusJSON(400, map[string]interface{}{"error": err.Error()})
		}).
		Complete())
	v1.DELETE("logoff", userController.Logoff)
}
