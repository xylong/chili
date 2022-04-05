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

func (c *Chili) Get(relativePath string, handler interface{}) *Chili {
	c.handle(http.MethodGet, relativePath, handler)
	return c
}

func (c *Chili) Post(relativePath string, handler interface{}) *Chili {
	c.handle(http.MethodPost, relativePath, handler)
	return c
}

func (c *Chili) Delete(relativePath string, handler interface{}) *Chili {
	c.handle(http.MethodDelete, relativePath, handler)
	return c
}

func (c *Chili) Patch(relativePath string, handler interface{}) *Chili {
	c.handle(http.MethodPatch, relativePath, handler)
	return c
}

func (c *Chili) Put(relativePath string, handler interface{}) *Chili {
	c.handle(http.MethodPut, relativePath, handler)
	return c
}

func (c *Chili) Options(relativePath string, handler interface{}) *Chili {
	c.handle(http.MethodOptions, relativePath, handler)
	return c
}

func (c *Chili) Any(relativePath string, handler interface{}) *Chili {
	c.Get(relativePath, handler)
	c.Post(relativePath, handler)
	c.Delete(relativePath, handler)
	c.Patch(relativePath, handler)
	c.Put(relativePath, handler)
	c.Options(relativePath, handler)
	c.handle(http.MethodHead, relativePath, handler)
	c.handle(http.MethodConnect, relativePath, handler)
	c.handle(http.MethodTrace, relativePath, handler)
	return c
}

func (c *Chili) handle(httpMethod, relativePath string, handlerFunc interface{}) *Chili {
	c.Engine.Handle(httpMethod, relativePath, getHandleFunc(handlerFunc))
	return c
}

// Launch 启动
func (c *Chili) Launch() {
	c.Run()
}
