package routes

import (
	"github.com/xylong/chili"
	v1 "github.com/xylong/chili/test/internal/api/v1"
	. "github.com/xylong/chili/test/internal/middleware"
)

// Version1 版本1
func Version1(group *chili.Group) {
	group.Post("register", v1.Register)
	group.Post("login", v1.Login)

	group.Group("", func(group *chili.Group) {
		group.Group("users", func(users *chili.Group) {
			users.Get("", v1.Me)
			users.Patch("", v1.Update)
			users.Delete("", v1.Logoff)
			users.Get(":id", v1.Show)

			users.Group("articles", func(articles *chili.Group) {
				articles.Post("", v1.Store)
			}, NewPermission())
		})

	}, NewAuthorization())
}
