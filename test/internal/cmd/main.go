package main

import (
	"github.com/xylong/chili"
	v2 "github.com/xylong/chili/test/internal/api/v2"
	"github.com/xylong/chili/test/internal/middleware"
	"github.com/xylong/chili/test/internal/routes"
)

func main() {
	chili.Ignite().
		Group("v1", routes.Version1, middleware.NewCsrf()).
		Group("v2", func(group *chili.Group) {
			group.Delete("logoff", v2.Logoff)
		}, middleware.NewCsrf()).
		Launch()
}
