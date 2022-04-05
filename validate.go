package chili

import "github.com/go-playground/validator/v10"

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
