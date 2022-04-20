package chili

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Group 路由组
type Group struct {
	*gin.RouterGroup
	group string
	middlewares
}

func NewGroup(routerGroup *gin.RouterGroup) *Group {
	return &Group{RouterGroup: routerGroup}
}

// Group 路由分组，通过闭包处理路由前缀和中间件
func (g Group) Group(group string, callback func(*Group), middlewares ...middleware) {
	g.middlewares = append(g.middlewares, middlewares...)
	g.group += group + "/"
	callback(&g)
}

// POST post请求
func (g *Group) POST(relativePath string, handler interface{}) *Group {
	g.handle(http.MethodPost, relativePath, handler)
	return g
}

// handle 注册路由处理函数
func (g *Group) handle(httpMethod, relativePath string, handler interface{}) *Group {
	if h := convert(handler); h != nil {
		g.Handle(httpMethod, relativePath, func(context *gin.Context) {
			ctx := &Context{context}
			g.middlewares.before(ctx)

			h(ctx)
		})
	}
	return g
}
