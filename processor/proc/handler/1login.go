package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

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
	url := cfg.AuthCodeURL(config.OauthStateString)
	return c.JSON(http.StatusOK, map[string]string{"redirectUrl": url})
}
