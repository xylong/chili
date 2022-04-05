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
	return chili
}

func (c *Chili) Post(relativePath string, handlerFunc gin.HandlerFunc) *Chili {
	c.Engine.POST(relativePath, handlerFunc)
	return c
}

func (c *Chili) Delete(relativePath string, handlerFunc gin.HandlerFunc) *Chili {
	c.Engine.DELETE(relativePath, handlerFunc)
	return c
}

// Launch 启动
func (c *Chili) Launch() {
	c.Run()
}
