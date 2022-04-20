package main

import (
	"github.com/xylong/chili"
	"github.com/xylong/chili/test/internal/middleware"
	"github.com/xylong/chili/test/internal/routes"
)

func main() {
	chili.Ignite().
		Group("v1", routes.Version1, middleware.NewCsrf()).
		Launch()
}
