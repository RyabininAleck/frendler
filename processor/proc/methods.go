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
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:3000"}, // Здесь можно указать конкретные домены вместо '*'
		AllowMethods:     []string{echo.GET, echo.PUT, echo.POST, echo.DELETE, echo.OPTIONS},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowCredentials: true,
	}))

	//todo установить TokenMiddleware для всех кроме login callback
	//e.Use(p.Handler.TokenMiddleware)
	//todo перенести "secret-key" в конфиг.
	//store := sessions.NewCookieStore([]byte("secret"))
	//store.Options = &sessions.Options{
	//	Path:     "/",
	//	MaxAge:   60,   // Время жизни куки в секундах (1 мин)
	//	HttpOnly: true, // HTTPOnly флаг
	//}
	//api.Use(session.Middleware(store))

	api := e.Group("/api")
	apiV1 := api.Group("/v1")
	p.RegisterHandlers(apiV1)

	// Запуск сервера
	return e.Start(":8080")
}

func (p *Processor) Stop() {

}
