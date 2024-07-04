package proc

import (
	"github.com/labstack/echo/v4"

	"frendler/processor/proc/handler"
)

func (p *Processor) RegisterHandlers(api *echo.Group) {

	api.GET("/hello", handler.HelloHandler)

	//todo проверить
	get := api.Group("/:id")
	get.GET("/profile", p.Handler.GetProfile)   // +
	get.GET("/settings", p.Handler.GetSettings) // +

	createUserApi := api.Group("/createUser")
	createUserApi.POST("/byNumber", p.Handler.CreateUserByNumber) // +
	createUserApi.POST("/byEmail", p.Handler.CreateUserByEmail)
	createUserApi.POST("/byVK", p.Handler.CreateUserByVk)

	addProfile := api.Group("/addProfile")
	addProfile.POST("/vk", p.Handler.AddVKProfile)
	//	addProfile.POST("/telegram", p.Handler.AddTelegramProfile)

	//addProfileCallback := addProfile.Group("/callback")
	//addProfileCallback.POST("/vk", p.Handler.AddVKProfileCallback)
	//addProfileCallback.POST("/telegram", p.Handler.AddTelegramProfileCallback)

	// todo объединение контакта, группы.

}
