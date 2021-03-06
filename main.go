package main

import (
	"github.com/labstack/echo"
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "swagtest/docs" // docs is generated by Swag CLI, you have to import it.
)

// @title Swagger Example API
// @version 1.0
// @description This is a test server
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1:1323
// @BasePath /
func main() {
	e := echo.New()

	e.GET("/hello-world", hello)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start("127.0.0.1:1323"))
}
