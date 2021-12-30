package handler

import (
	"github.com/developernaren/go-tests/user"
	"github.com/labstack/echo/v4"
	"net/http"
)

func Create(c echo.Context) error {

	name := c.FormValue("name")
	createdUser, err := user.Create(name)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Error creating user",
		})
	}

	return c.JSON(http.StatusOK, createdUser)
}
