package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// HelloHandler обработчик GET запроса "/"
func HelloHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
