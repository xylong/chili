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

func (c *Chili) Group() *Chili {
	return c
}

// Launch 启动
func (c *Chili) Launch() {
	c.Run()
}
