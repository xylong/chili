package chili

import "github.com/gin-gonic/gin"

// Group 路由组
type Group struct {
	*gin.RouterGroup
	group string
}

func NewGroup(routerGroup *gin.RouterGroup) *Group {
	return &Group{RouterGroup: routerGroup}
}

// Group 路由分组
func (g Group) Group(group string, callback func(group *Group)) {
	g.group += group + "/"
	callback(&g)
}
