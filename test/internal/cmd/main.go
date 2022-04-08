package main

import (
	"github.com/xylong/chili"
	v1 "github.com/xylong/chili/test/internal/api/v1"
	v2 "github.com/xylong/chili/test/internal/api/v2"
)

func main() {
	chili.Ignite().
		Group("v1", func(group *chili.Group) {
			group.Post("register", v1.Register)

			group.Group("a", func(group *chili.Group) {
				group.Post("login", v1.Login)

				group.Group("b", func(group *chili.Group) {
					group.Delete("logoff", v1.Logoff)
				})
			})
		}).
		Group("v2", func(group *chili.Group) {
			group.Delete("logoff", v2.Logoff)
		}).
		Launch()
}
