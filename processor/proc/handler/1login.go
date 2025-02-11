package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"

	"frendler/processor/config"
	googleModels "frendler/processor/models/google"
)

func (h *HandlerImpl) LoginByGoogle(c echo.Context) error {

	req := new(googleModels.LoginRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	cfg := config.GoogleOauth
	cfg.RedirectURL = req.RedirectURL
	url := cfg.AuthCodeURL(config.OauthStateString, oauth2.AccessTypeOffline)
	return c.JSON(http.StatusOK, map[string]string{"redirectUrl": url})
}

func (h *HandlerImpl) LoginByVK(c echo.Context) error {
	return c.NoContent(http.StatusNotImplemented)
}
func (h *HandlerImpl) LoginByTg(c echo.Context) error {
	return c.NoContent(http.StatusNotImplemented)
}

func (h *HandlerImpl) LoginByWhatsUp(c echo.Context) error {
	return c.NoContent(http.StatusNotImplemented)
}
