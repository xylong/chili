package chili

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
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
	g.group += fmt.Sprintf("/%s", strings.Trim(group, "/"))
	callback(&g)
}

// Get get请求
func (g *Group) Get(relativePath string, handler interface{}) *Group {
	g.handle(http.MethodGet, relativePath, handler)
	return g
}

// Post post请求
func (g *Group) Post(relativePath string, handler interface{}) *Group {
	g.handle(http.MethodPost, relativePath, handler)
	return g
}

// Patch patch请求
func (g *Group) Patch(relativePath string, handler interface{}) *Group {
	g.handle(http.MethodPost, relativePath, handler)
	return g
}

// Delete delete请求
func (g *Group) Delete(relativePath string, handler interface{}) *Group {
	g.handle(http.MethodDelete, relativePath, handler)
	return g
}

// handle 注册路由处理函数
func (g *Group) handle(httpMethod, relativePath string, handler interface{}) *Group {
	if value, ok := isAction(handler); ok {
		// 判断是否为gin原生的HandlerFunc
		g.Handle(httpMethod, fmt.Sprintf("%s/%s", g.group, relativePath), func(context *gin.Context) {
			ctx := &Context{context}
			g.middlewares.before(ctx)

			if value.Type().NumIn() == oneParam {
				value.Interface().(func(*Context))(ctx)
			} else {
				convert(ctx, value)
			}
		})
	}

	return g
}
