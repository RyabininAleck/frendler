package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *HandlerImpl) TokenMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sessionToken, sessionTokenErr := c.Cookie("sessionToken")
		userID, userIDErr := c.Cookie("userId")
		if userIDErr != nil || sessionTokenErr != nil {
			return c.NoContent(http.StatusUnauthorized)
		}

		token := h.DB.CheckToken(userID.Value, sessionToken.Value)
		if token == nil || !token.IsActive {
			return c.NoContent(http.StatusUnauthorized)
		}

		return next(c)
	}
}
