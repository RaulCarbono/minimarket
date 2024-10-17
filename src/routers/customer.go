package routers

import (
	"github.com/go/mini_market/src/handlers"
	"github.com/go/mini_market/src/server"
	"github.com/labstack/echo/v4"
)

func CustomerRouter(s server.Server, e *echo.Echo) {
	customerGroup := e.Group("/customer")
	customerGroup.GET("/info/:id", handlers.GetCustomerByIdHandler(s))
}
