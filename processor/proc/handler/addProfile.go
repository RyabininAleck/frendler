package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *HandlerImpl) AddVKProfile(c echo.Context) error {
	return c.String(http.StatusOK, "AddVKProfile")
}

func (h *HandlerImpl) AddTelegramProfile(c echo.Context) error {
	return c.String(http.StatusOK, "AddTelegramProfile!")
}
