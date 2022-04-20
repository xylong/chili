package chili

import (
	"github.com/gin-gonic/gin"
)

// Chili 脚手架核心
type Chili struct {
	*gin.Engine
}

// Ignite 初始化
func Ignite() *Chili {
	chili := &Chili{Engine: gin.New()}
	chili.Use(gin.Logger(), exception())

	return chili
}

// Group 路由分组
func (c *Chili) Group(group string, callback func(group *Group), middlewares ...middleware) *Chili {
	g := NewGroup(c.Engine.Group(group))
	g.middlewares = append(g.middlewares, middlewares...)
	callback(g)

	return c
}

// Launch 启动
func (c *Chili) Launch() {
	c.Run()
}
