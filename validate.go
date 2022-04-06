package chili

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var v *validator.Validate

func init() {
	v = validator.New()
}

func Validate(param interface{}) (bool, string) {
	err := v.Struct(param)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		if len(errs) > 0 {
			err := errs[0]
			return false, err.Error()
		}
	}

	return true, ""
}

// Param 参数
type Param interface {
	Bind(*gin.Context) error
}
