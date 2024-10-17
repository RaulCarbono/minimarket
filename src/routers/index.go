package routers

import (
	"github.com/go/mini_market/src/middleware"
	"github.com/go/mini_market/src/server"
	"github.com/labstack/echo/v4"
)

func BindRoutes(s server.Server, e *echo.Echo) {
	e.Use(middleware.CheckAuthMiddleware(s))
	HomeRouter(s, e)
	UserRouter(s, e)
	CustomerRouter(s, e)
	AuthenticateRouter(s, e)
}
