package routers

import (
	"github.com/go/mini_market/src/handlers"
	"github.com/go/mini_market/src/server"
	"github.com/labstack/echo/v4"
)

func ProductRouter(s server.Server, e *echo.Echo) {
	productGroup := e.Group("/product")
	productGroup.GET("/info/:id", handlers.GetProductByIdHandler(s))
	productGroup.POST("", handlers.InsertProductHandler(s))
	productGroup.GET("/info", handlers.GetProductdHandler(s))

}
