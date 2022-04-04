package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/xylong/chili/test/internal/request"
)

var v *validator.Validate

func init() {
	v = validator.New()
}

func Validate(param interface{}) (bool, string) {
	err := v.Struct(param)
	errs := err.(validator.ValidationErrors)
	if len(errs) > 0 {
		err := errs[0]
		return false, err.Error()
	}
	return true, ""
}

func Login(ctx *gin.Context) {
	param := request.LoginParam{}
	err := ctx.ShouldBind(&param)
	if err != nil {
		ctx.String(400, "请求参数异常")
		return
	}
	ctx.String(200, "登录")
}
