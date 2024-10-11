package routers

import (
	"github.com/go/mini_market/src/handlers"
	"github.com/go/mini_market/src/server"
	"github.com/labstack/echo/v4"
)

func AuthenticateRouter(s server.Server, r *echo.Router) {
	r.Add("POST", "/auth/signup", handlers.SignUpHandler(s))
	r.Add("POST", "/auth/login", handlers.LoginHandler(s))
}
