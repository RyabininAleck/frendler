package proc

import (
	"github.com/labstack/echo/v4"

	"frendler/processor/proc/handler"
)

func RegisterHandlers(api *echo.Group) {

	api.GET("/hello", handler.HelloHandler)

	get := api.Group("/:id")
	get.GET("/profile", handler.GetProfile)
	get.GET("/settings", handler.GetSettings)

	createUserApi := api.Group("/createUser")
	createUserApi.POST("/byEmail", handler.CreateUserByEmail)
	createUserApi.POST("/byNumber", handler.CreateUserByNumber)
	createUserApi.POST("/byVK", handler.CreateUserByVk)

	addProfile := api.Group("/addProfile")
	addProfile.POST("/vk", handler.AddVKProfile)
	addProfile.POST("/telegram", handler.AddTelegramProfile)

}
