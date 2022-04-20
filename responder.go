package chili

import (
	"net/http"
	"reflect"
	"sync"
)

var (
	once          sync.Once
	responderList []Responder
)

type Responder interface {
	Return() func(*Context)
}

type StringResponder func(*Context, Param) string

func (r StringResponder) Return() func(*Context) {
	return func(context *Context) {
		context.String(http.StatusOK, r(context, nil))
	}
}

func getResponderList() []Responder {
	once.Do(func() {
		responderList = []Responder{
			(StringResponder)(nil),
		}
	})

	return responderList
}

type (
	handlerFunc func(...Param) func(output)
	output      func(*Context, interface{})
)

func wrap(ctx *Context) handlerFunc {
	return func(param ...Param) func(output) {
		// todo 参数绑定验证

		// todo 结果处理
		return func(o output) {
			o(ctx, "a")
		}
	}
}

func convert(handler interface{}) func(*Context) {
	value := reflect.ValueOf(handler)

	for _, responder := range getResponderList() {
		r := reflect.TypeOf(responder)
		if value.Type().ConvertibleTo(r) {
			return value.Convert(r).Interface().(Responder).Return()
		}
	}

	return nil
}
