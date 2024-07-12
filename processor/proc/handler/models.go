package handler

import (
	"github.com/labstack/echo/v4"

	"frendler/processor/db"
)

type HandlerImpl struct {
	DB db.DB
}

type Handler interface {
	LoginByGoogle(c echo.Context) error
	HandleGoogleCallback(c echo.Context) error

	//
	//CreateUserByNumber(c echo.Context) error
	//CreateUserByVk(c echo.Context) error
	//
	//GetProfile(c echo.Context) error
	GetSettings(c echo.Context) error
	//
	//AddVKProfile(c echo.Context) error
	//AddTelegramProfile(c echo.Context) error
	TokenMiddleware(next echo.HandlerFunc) echo.HandlerFunc
}
