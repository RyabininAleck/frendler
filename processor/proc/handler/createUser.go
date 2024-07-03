package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateUserByEmail(c echo.Context) error {
	return c.String(http.StatusOK, "CreateUserByEmail")
}

func CreateUserByNumber(c echo.Context) error {
	return c.String(http.StatusOK, "CreateUserByNumber")
}

func CreateUserByVk(c echo.Context) error {
	return c.String(http.StatusOK, "CreateUserByVk!")
}

func CreateUser() {

}
