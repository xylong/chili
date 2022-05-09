package chili

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
)

var responderList []Responder

func init() {
	responderList = []Responder{
		new(stringResponder),
		new(mapResponder),
	}
}

// Responder 响应器
type Responder interface {
	respond() gin.HandlerFunc
}

type (
	stringResponder func(*Context) string
	mapResponder    func(*Context) map[string]interface{}
)

func (r stringResponder) respond() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.String(http.StatusOK, r(NewContext(context)))
	}
}

func (r mapResponder) respond() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(http.StatusOK, r(NewContext(context)))
	}
}

func convert(handler interface{}) gin.HandlerFunc {
	value := reflect.ValueOf(handler)

	for _, responder := range responderList {
		val := reflect.ValueOf(responder).Elem()
		if value.Type().ConvertibleTo(val.Type()) {
			val.Set(value)
			return val.Interface().(Responder).respond()
		}
	}

	return nil
}
