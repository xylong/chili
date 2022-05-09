package chili

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
)

var ResponderList []Responder

func init() {
	ResponderList = []Responder{
		new(StringResponder),
	}
}

// Responder 响应器
type Responder interface {
	Return() gin.HandlerFunc
}

type (
	StringResponder func(*gin.Context) string
	JsonResponder   func(*gin.Context) gin.H
)

func (r StringResponder) Return() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.String(http.StatusOK, r(context))
	}
}

func convert(handler interface{}) gin.HandlerFunc {
	value := reflect.ValueOf(handler)

	for _, responder := range ResponderList {
		val := reflect.ValueOf(responder).Elem()
		if value.Type().ConvertibleTo(val.Type()) {
			val.Set(value)
			return val.Interface().(Responder).Return()
		}
	}

	return nil
}
