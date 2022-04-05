package chili

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Chili 脚手架核心
type Chili struct {
	*gin.Engine
}

// Ignite 初始化
func Ignite() *Chili {
	chili := &Chili{Engine: gin.New()}
	return chili
}

// Get get请求
func (c *Chili) Get(relativePath string, handler interface{}) *Chili {
	c.handle(http.MethodGet, relativePath, handler)
	return c
}

// Post post请求
func (c *Chili) Post(relativePath string, handler interface{}) *Chili {
	c.handle(http.MethodPost, relativePath, handler)
	return c
}

// Delete delete请求
func (c *Chili) Delete(relativePath string, handler interface{}) *Chili {
	c.handle(http.MethodDelete, relativePath, handler)
	return c
}

// Patch patch请求
func (c *Chili) Patch(relativePath string, handler interface{}) *Chili {
	c.handle(http.MethodPatch, relativePath, handler)
	return c
}

// Put put请求
func (c *Chili) Put(relativePath string, handler interface{}) *Chili {
	c.handle(http.MethodPut, relativePath, handler)
	return c
}

// Options options请求
func (c *Chili) Options(relativePath string, handler interface{}) *Chili {
	c.handle(http.MethodOptions, relativePath, handler)
	return c
}

// Any 任意类型请求
func (c *Chili) Any(relativePath string, handler interface{}) *Chili {
	c.handle(http.MethodGet, relativePath, handler)
	c.handle(http.MethodPost, relativePath, handler)
	c.handle(http.MethodDelete, relativePath, handler)
	c.handle(http.MethodPatch, relativePath, handler)
	c.handle(http.MethodPut, relativePath, handler)
	c.handle(http.MethodOptions, relativePath, handler)
	c.handle(http.MethodHead, relativePath, handler)
	c.handle(http.MethodConnect, relativePath, handler)
	c.handle(http.MethodTrace, relativePath, handler)
	return c
}

// handle 注册路由函数
func (c *Chili) handle(httpMethod, relativePath string, handler interface{}) *Chili {
	if value, ok := isAction(handler); ok {
		// 如果只有1个参数说明是gin原生的HandlerFunc
		if value.Type().NumIn() == oneParam {
			c.Engine.Handle(httpMethod, relativePath, value.Interface().(func(ctx *gin.Context)))
		} else {
			c.Engine.Handle(httpMethod, relativePath, func(context *gin.Context) {
				proxyHandlerFunc(context, value)
			})
		}
	}

	return c
}

// Launch 启动
func (c *Chili) Launch() {
	c.Run()
}
