package chili

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// Group 路由组
type Group struct {
	*gin.RouterGroup
	group string
}

func NewGroup(routerGroup *gin.RouterGroup) *Group {
	return &Group{RouterGroup: routerGroup}
}

// Group 路由分组
func (g Group) Group(group string, callback func(*Group)) {
	g.group += group + "/"
	callback(&g)
}

func (g *Group) GET(relativePath string, handler interface{}) {
	if relativePath != "/*any" {
		g.handle(http.MethodGet, relativePath, handler)
	} else {
		g.Handle(http.MethodGet, relativePath, handler.(gin.HandlerFunc))
	}
}

func (g *Group) POST(relativePath string, handler interface{}) {
	g.handle(http.MethodPost, relativePath, handler)
}

func (g *Group) DELETE(relativePath string, handler interface{}) {
	g.handle(http.MethodDelete, relativePath, handler)
}

func (g *Group) handle(httpMethod, relativePath string, handler interface{}) {
	if f := convert(handler); f != nil {
		g.Handle(httpMethod, strings.TrimRight(g.group+"/"+relativePath, "/"), f)
	}
}
