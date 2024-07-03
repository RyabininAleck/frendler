package proc

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (p *Processor) RunTasks() {

}

func (p *Processor) Run() error {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	api := e.Group("/api")
	apiV1 := api.Group("/v1")
	RegisterHandlers(apiV1)

	// Запуск сервера
	return e.Start(":8080")
}

func (p *Processor) Stop() {

}
