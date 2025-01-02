package routers

import (
	"github.com/go/mini_market/src/handlers"
	"github.com/go/mini_market/src/server"
	"github.com/labstack/echo/v4"
)

func AuthenticateRouter(s server.Server, e *echo.Echo) {
	authGroup := e.Group("/auth")
	authGroup.POST("/signup", handlers.SignUpHandler(s))
	authGroup.POST("/login", handlers.LoginHandler(s))
}
