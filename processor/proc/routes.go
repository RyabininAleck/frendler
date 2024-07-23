package proc

import (
	"github.com/labstack/echo/v4"

	"frendler/processor/proc/handler"
)

func (p *Processor) RegisterHandlers(api *echo.Group) {

	api.GET("/hello", handler.HelloHandler)

	createUserApi := api.Group("/login")
	createUserApi.POST("/byGoogle", p.Handler.LoginByGoogle)

	CallbackApi := api.Group("/callback")
	CallbackApi.GET("/google", p.Handler.HandleGoogleCallback)

	userApi := api.Group("/user")
	userApi.Use(p.Handler.TokenMiddleware)

	userApi.GET("/settings", p.Handler.GetSettings)
	userApi.GET("/contactStats", p.Handler.GetContactStats)
	userApi.GET("/qr", p.Handler.GetQRCode)

	userUpdate := userApi.Group("/update")
	userUpdate.GET("/google", p.Handler.GoogleContactUpdate)

	////todo проверить
	//get := api.Group("/:id")
	//get.GET("/profile", p.Handler.GetProfile)   // +
	//get.GET("/settings", p.Handler.GetSettings) // +

	//createUserApi := api.Group("/login")
	//createUserApi.POST("/byNumber", p.Handler.CreateUserByNumber) // +
	//createUserApi.POST("/byGoogle", p.Handler.LoginByGoogle)
	//createUserApi.POST("/byVK", p.Handler.CreateUserByVk)

	//add := api.Group("/:id")
	//addProfile := add.Group("/addProfile")
	//addProfile.POST("/vk", p.Handler.AddVKProfile)
	//	addProfile.POST("/telegram", p.Handler.AddTelegramProfile)

	//addProfileCallback := addProfile.Group("/callback")
	//addProfileCallback.POST("/vk", p.Handler.AddVKProfileCallback)
	//addProfileCallback.POST("/telegram", p.Handler.AddTelegramProfileCallback)

	//addFriends := add.Group("/addFriends")
	//addFriends.POST("/vk", p.Handler.AddVkFriends)
	//addFriends.POST("/telegram", p.Handler.AddTelegramFriends)

	// todo объединение контакта, группы.

}
