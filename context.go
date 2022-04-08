package chili

import (
	"github.com/gin-gonic/gin"
	"strings"
)

// Context 上下文
type Context struct {
	*gin.Context
}

// Domain 获取域名
func (c *Context) Domain() string {
	return c.Request.Host[:strings.Index(c.Request.Host, ":")]
}
