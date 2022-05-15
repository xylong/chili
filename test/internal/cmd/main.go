package main

import (
	"github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/xylong/chili"
	_ "github.com/xylong/chili/docs"
	"github.com/xylong/chili/test/internal/router"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      127.0.0.1:8080
// @BasePath  /v1

//go:generate swag init --parseDependency  --parseDepth=6 -g test/internal/cmd/main.go
func main() {
	chili.Ignite().
		Group("swagger", func(swagger *chili.Group) {
			swagger.GET("/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		}).
		Group("v1", router.V1).
		Launch()
}
