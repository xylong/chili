package chili

import (
	"github.com/gin-gonic/gin"
	"strings"
)

// Context 上下文
type Context struct {
	*gin.Context
}

func NewContext(context *gin.Context) *Context {
	return &Context{Context: context}
}

// Domain 域名
func (c *Context) Domain() string {
	return c.Request.Host[:strings.Index(c.Request.Host, ":")]
}
