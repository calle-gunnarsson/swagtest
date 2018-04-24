package main

import (
	"net/http"

	"github.com/labstack/echo"
)

// hello godoc
// @Summary Return "Hello World!"
// @Description return "Hello World!"
// @Tags HelloWorld
// @Accept  json
// @Produce  json
// @Success 200 {object} HelloWorld
// @Failure 500 {object} error
// @Router /hello-world [get]
func hello(c echo.Context) error {
	h := HelloWorld{"Hello World!"}
	return c.JSON(http.StatusOK, h)
}
