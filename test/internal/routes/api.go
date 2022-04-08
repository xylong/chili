package routes

import (
	"github.com/xylong/chili"
	v1 "github.com/xylong/chili/test/internal/api/v1"
)

// Version1 版本1
func Version1(group *chili.Group) {
	group.Post("register", v1.Register)

	group.Group("a", func(group *chili.Group) {
		group.Post("login", v1.Login)

		group.Group("b", func(group *chili.Group) {
			group.Delete("logoff", v1.Logoff)
		})
	})
}
