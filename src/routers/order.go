package routers

import (
	"github.com/go/mini_market/src/handlers"
	"github.com/go/mini_market/src/server"
	"github.com/labstack/echo/v4"
)

func OrderRouter(s server.Server, e *echo.Echo) {
	authGroup := e.Group("/order")
	authGroup.POST("", handlers.InsertOrderHandler(s))
	authGroup.POST("/add-item", handlers.AddItemHandler(s))
}
