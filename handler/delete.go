package handler

import (
	"fmt"
	"github.com/developernaren/go-tests/user"
	"github.com/labstack/echo/v4"
	"net/http"
)

func Delete(c echo.Context) error {

	userID := c.Param("id")
	done := user.DeleteByID(userID)

	if done {
		return c.JSON(http.StatusNoContent, nil)
	}

	return c.JSON(http.StatusInternalServerError, map[string]string{
		"message": fmt.Sprintf("User with id:%s could not be deleted!", userID),
	})
}
