package routers

import (
	"github.com/go/mini_market/src/handlers"
	"github.com/go/mini_market/src/server"
	"github.com/labstack/echo/v4"
)

func HomeRouter(s server.Server, r *echo.Router) {
	r.Add("GET", "/", handlers.HomeHandler(s))
}
