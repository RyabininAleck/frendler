package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func HelloHandler(c echo.Context) error {
	cookies := c.Cookies()
	response := make(map[string]string)
	for _, cookie := range cookies {
		response[cookie.Name] = cookie.Value
	}
	return c.JSON(http.StatusOK, response)
}
