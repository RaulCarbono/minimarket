package routers

import (
	"github.com/go/mini_market/src/handlers"
	"github.com/go/mini_market/src/server"
	"github.com/labstack/echo/v4"
)

func HomeRouter(s server.Server, e *echo.Echo) {
	HomeGroup := e.Group("/home")
	HomeGroup.GET("", handlers.HomeHandler(s))
}
