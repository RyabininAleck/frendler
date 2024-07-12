package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *HandlerImpl) TokenMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		idCookie, err := c.Cookie("id")
		if err != nil || idCookie.Value == "" {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"error": "Missing or invalid id cookie",
			})
		}

		tokenCookie, err := c.Cookie("token")
		if err != nil || tokenCookie.Value == "" {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"error": "Missing or invalid token cookie",
			})
		}

		ok, err := h.DB.CheckToken(idCookie.Value, tokenCookie.Value)
		if err != nil {
			return err
		}

		if !ok {
			return c.JSON(http.StatusTemporaryRedirect, "http://localhost:3000/login")
		}

		// Передаем запрос дальше
		return next(c)
	}
}
