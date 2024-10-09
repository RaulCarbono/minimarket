package routers

import (
	"github.com/go/mini_market/src/server"
	"github.com/labstack/echo/v4"
)

func BindRoutes(s server.Server, r *echo.Router) {
	HomeRouter(s, r)
	UserRouter(s, r)
	AuthenticateRouter(s, r)
}
