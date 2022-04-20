package routes

import (
	"github.com/xylong/chili"
	v1 "github.com/xylong/chili/test/internal/api/v1"
)

// Version1 版本1
func Version1(group *chili.Group) {
	group.POST("register", v1.Register)
}
