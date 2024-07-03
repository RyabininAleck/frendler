package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetProfile(c echo.Context) error {
	return c.String(http.StatusOK, "GetProfile")
}

func GetSettings(c echo.Context) error {
	return c.String(http.StatusOK, "GetSettings")
}
