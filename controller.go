package chili

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
)

const (
	// 上下文类型
	contextType = "gin.Context"
)

// proxyHandlerFunc 根据函数参数实现自动创建
// 获取函数的第二个参数Type，然后再通过反射创建实例
func proxyHandlerFunc(ctx *gin.Context, value reflect.Value) {
	// 获取第二个参数的类型，并创建实例
	paramType := value.Type().In(1).Elem()
	param := reflect.New(paramType).Interface()

	// 参数绑定
	err := ctx.ShouldBind(param)
	if err != nil {
		fmt.Println(err)
		ctx.String(http.StatusBadRequest, "请求参数异常")
		return
	}

	// 参数验证
	ok, msg := Validate(param)
	if !ok {
		ctx.String(http.StatusBadRequest, msg)
		return
	}

	// 调用真实handleFunc
	value.Call(valOf(ctx, param))
}

func valOf(i ...interface{}) (rv []reflect.Value) {
	for _, i2 := range i {
		rv = append(rv, reflect.ValueOf(i2))
	}

	return
}

// convert 将路由函数转为gin.HandlerFunc
func convert(value reflect.Value) func(ctx *gin.Context) {
	paramNum := value.Type().NumIn()
	return func(ctx *gin.Context) {
		// 如果只有1个参数说明是gin原生的HandlerFunc
		if paramNum == 1 {
			value.Call(valOf(ctx))
			return
		}

		proxyHandlerFunc(ctx, value)
	}
}

// isAction 判断是否为控制器方法
func isAction(handler interface{}) (value reflect.Value, ok bool) {
	funcType := reflect.TypeOf(handler)

	if funcType.Kind() != reflect.Func {
		return
	}

	// 判断第一个参数是否为*gin.Context
	paramNum := funcType.NumIn()
	if paramNum == 0 {
		return
	}

	if t := funcType.In(0); t.Kind() == reflect.Ptr && t.Elem().String() == contextType {
		value, ok = reflect.ValueOf(handler), true
	}

	return
}
