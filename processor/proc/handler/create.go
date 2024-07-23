package handler

import (
	"net/http"
	"strconv"

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

func NotionHandler(c echo.Context) error {
	cookieId, err := c.Cookie("userId")
	userId, err := strconv.Atoi(cookieId.Value)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"userId": userId})
}
