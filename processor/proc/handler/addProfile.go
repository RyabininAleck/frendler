package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func AddVKProfile(c echo.Context) error {
	return c.String(http.StatusOK, "AddVKProfile")
}

func AddTelegramProfile(c echo.Context) error {
	return c.String(http.StatusOK, "AddTelegramProfile!")
}
