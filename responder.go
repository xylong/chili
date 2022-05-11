package chili

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
)

var responderList []Responder

func init() {
	responderList = []Responder{
		(stringResponder)(nil),
		(jsonResponder)(nil),
		(defaultResponder)(nil),
	}
}

// Responder 响应器
type Responder interface {
	respond() gin.HandlerFunc
}

type stringResponder func(*Context) string

func (r stringResponder) respond() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.String(http.StatusOK, r(NewContext(context)))
	}
}

type (
	Json          interface{}
	jsonResponder func(*Context) Json
)

func (r jsonResponder) respond() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(http.StatusOK, r(NewContext(context)))
	}
}

type defaultResponder func(*Context)

func (r defaultResponder) respond() gin.HandlerFunc {
	return func(context *gin.Context) {
		r(NewContext(context))
	}
}

func convert(handler interface{}) gin.HandlerFunc {
	value := reflect.ValueOf(handler)

	for _, responder := range responderList {
		t := reflect.TypeOf(responder)
		if value.Type().ConvertibleTo(t) {
			return value.Convert(t).Interface().(Responder).respond()
		}
	}

	return nil
}
