package chili

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
)

const (
	// 参数个数
	oneParam = iota + 1
	twoParam

	// 上下文类型
	contextType = "gin.Context"
)

// proxyHandlerFunc 根据函数参数实现自动创建
func proxyHandlerFunc(ctx *gin.Context, value reflect.Value) {
	// 获取第二个参数的类型，并创建实例
	paramType := value.Type().In(twoParam - 1).Elem()
	param := reflect.New(paramType).Interface()

	// 参数绑定
	var err error
	if p, ok := param.(Param); ok {
		err = p.Bind(ctx)
	} else {
		err = ctx.ShouldBind(param)
	}

	if err != nil {
		fmt.Println(err)
		ctx.String(http.StatusBadRequest, "请求参数异常")
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

// isAction 判断是否为控制器方法
func isAction(handler interface{}) (reflect.Value, bool) {
	funcType, funcValue := reflect.TypeOf(handler), reflect.ValueOf(handler)

	if funcType.Kind() != reflect.Func || funcType.NumIn() < oneParam {
		return funcValue, false
	}

	p := funcType.In(0)
	return funcValue, p.Kind() == reflect.Ptr && p.Elem().String() == contextType
}
