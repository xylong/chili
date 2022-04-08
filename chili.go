package chili

import (
	"github.com/gin-gonic/gin"
)

// Chili 脚手架核心
type Chili struct {
	*gin.Engine
	group *Group
}

// Ignite 初始化
func Ignite() *Chili {
	chili := &Chili{Engine: gin.New()}
	chili.Use(gin.Logger(), gin.Recovery())

	return chili
}

// Group 路由分组
func (c *Chili) Group(group string, callback func(group *Group)) *Chili {
	c.group = NewGroup(c.Engine.Group(group))
	callback(c.group)
	return c
}

// Launch 启动
func (c *Chili) Launch() {
	c.Run()
}
