package chili

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Group 路由组
type Group struct {
	*gin.RouterGroup
	group string
}

type handlerFunc func(*gin.Context) *Response

func NewGroup(routerGroup *gin.RouterGroup) *Group {
	return &Group{RouterGroup: routerGroup}
}

// Group 路由分组
func (g Group) Group(group string, callback func(group *Group)) {
	g.group += group + "/"
	callback(&g)
}

func (g *Group) handle(httpMethod, relativePath string, handler interface{}) {
	switch h := handler.(type) {
	case func(*gin.Context):
		g.Handle(httpMethod, relativePath, h)
	case func(*gin.Context) *Response:
		g.Handle(httpMethod, relativePath, convert(h))
	}
}

func (g *Group) GET(relativePath string, handler interface{}) {
	g.handle(http.MethodGet, relativePath, handler)
}

func (g *Group) POST(relativePath string, handler interface{}) {
	g.handle(http.MethodPost, relativePath, handler)
}

func convert(handler handlerFunc) gin.HandlerFunc {
	return func(context *gin.Context) {
		f := handler(context)

		switch t := f.Data().(type) {
		case string:
			context.String(http.StatusOK, t)
		case gin.H:
			context.JSONP(http.StatusOK, t)
		}
	}
}
