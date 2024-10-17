package routers

import (
	"github.com/go/mini_market/src/handlers"
	"github.com/go/mini_market/src/server"
	"github.com/labstack/echo/v4"
)

func UserRouter(s server.Server, e *echo.Echo) {
	userGroup := e.Group("/user")
	userGroup.GET("/info/:id", handlers.GetUserByIdHandler(s))
	userGroup.GET("/info", handlers.GetUserHandler(s))

}
