package main

import (
	"github.com/developernaren/go-tests/handler"
	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()

	e.POST("/users", handler.Create)
	e.GET("/users", handler.List)
	e.GET("/users/:id", handler.Read)
	e.PUT("/users/:id", handler.Update)
	e.DELETE("users/:id", handler.Delete)

	e.Logger.Fatal(e.Start(":1323"))
}
