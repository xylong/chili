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

type StringResponder func(*Context) string

func (r StringResponder) Return() func(*Context) {
	return func(context *Context) {
		context.String(http.StatusOK, r(context))
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
