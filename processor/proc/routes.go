package proc

import (
	"github.com/labstack/echo/v4"

	"frendler/processor/proc/handler"
)

func (p *Processor) RegisterHandlers(api *echo.Group) {

	api.GET("/hello", handler.HelloHandler)

	createUserApi := api.Group("/login")
	createUserApi.POST("/byGoogle", p.Handler.LoginByGoogle)
	//todo createUserApi.POST("/byVk", p.Handler.LoginByVk)

	CallbackApi := api.Group("/callback")
	CallbackApi.GET("/google", p.Handler.HandleGoogleCallback)
	//todo CallbackApi.GET("/vk", p.Handler.HandleVkCallback)

	userApi := api.Group("/user")
	userApi.Use(p.Handler.TokenMiddleware)

	userApi.GET("/settings", p.Handler.GetSettings)
	userApi.GET("/contactStats", p.Handler.GetContactStats)
	userApi.GET("/qr", p.Handler.GetQRCode)
	//todo userApi.GET("/profile", p.Handler.GetProfile)

	userUpdate := userApi.Group("/update")
	userUpdate.GET("/google", p.Handler.GoogleContactUpdate)

}
