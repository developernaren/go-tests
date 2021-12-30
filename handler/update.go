package handler

import (
	"github.com/developernaren/go-tests/user"
	"github.com/labstack/echo/v4"
	"net/http"
)

func Update(c echo.Context) error {

	updatedUser := user.New(c.Param("id"), c.FormValue("name"))
	done := user.Update(updatedUser)

	if done {
		return c.JSON(http.StatusNoContent, updatedUser)
	}

	return c.JSON(http.StatusInternalServerError, map[string]string{
		"message": "User Could not be updated",
	})
}