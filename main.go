package main

import (
	"go_bedu/config"
	_ "go_bedu/docs" // docs is generated by Swag CLI, you have to import it.
	"go_bedu/helpers"
	m "go_bedu/middlewares"
	"go_bedu/routes"

	"github.com/labstack/echo/v4"
	mid "github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title           bEDU Documentation API
// @version         1.0
// @termsOfService  http://swagger.io/terms/

// @contact.name   Dicoding Cycle 4 TEAM C23-M4058
// @contact.url    https://github.com/Capstone-Dicoding-Cycle-4-C23-M4058

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// capstone.keyzex.com
// localhost:8080

// @host      capstone.keyzex.com
// @BasePath  /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	e := echo.New()

	m.Log(e)
	e.Pre(mid.RemoveTrailingSlash())

	db, err := config.ConnectDB()
	if err != nil {
		panic(err)
	}

	err = config.MigrateDB(db)
	if err != nil {
		panic(err)
	}

	routes.NewRoute(e, db)

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	var port = helpers.EnvPortOr("3000")

	// Start server with TLS
	e.Logger.Fatal(e.Start(port))
}
