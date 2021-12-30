package handler

import (
	"github.com/developernaren/go-tests/user"
	"github.com/labstack/echo/v4"
	"net/http"
)

func List(c echo.Context) error {

	return c.JSON(http.StatusOK, user.GetAll())
}
