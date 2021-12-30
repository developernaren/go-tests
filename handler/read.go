package handler

import (
	"fmt"
	"github.com/developernaren/go-tests/user"
	"github.com/labstack/echo/v4"
	"net/http"
)

func Read(c echo.Context) error {

	userID := c.Param("id")
	readUser := user.FindByID(userID)

	if readUser != nil {
		return c.JSON(http.StatusOK, readUser)
	}

	return c.JSON(http.StatusNotFound, map[string]string{
		"message": fmt.Sprintf("User with id:%s does not exist", userID),
	})
}
