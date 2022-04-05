package chili

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
)

// proxyHandlerFunc 根据函数参数实现自动创建
// 通过反射 handlerFunc 然后获取函数的第二个参数即可拿到 Type 然后再通过反射创建实例
func proxyHandlerFunc(ctx *gin.Context, handleFunc interface{}) {
	funcType := reflect.TypeOf(handleFunc)

	// 判断是否是func
	if funcType.Kind() != reflect.Func {
		panic("the route handlerFunc must be a function")
	}

	// 获取第二个参数的类型，并创建实例
	paramType := funcType.In(1).Elem()
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
	reflect.ValueOf(handleFunc).Call(valOf(ctx, param))
}

func valOf(i ...interface{}) (rv []reflect.Value) {
	for _, i2 := range i {
		rv = append(rv, reflect.ValueOf(i2))
	}

	return
}

func getHandleFunc(handler interface{}) func(ctx *gin.Context) {
	paramNum := reflect.TypeOf(handler).NumIn()
	return func(ctx *gin.Context) {
		// 如果只有1个参数说明是gin原生的HandlerFunc
		if paramNum == 1 {
			reflect.ValueOf(handler).Call(valOf(ctx))
			return
		}

		proxyHandlerFunc(ctx, handler)
	}
}
