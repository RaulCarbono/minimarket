package handlers

import (
	"github.com/go/mini_market/src/model"
	"github.com/go/mini_market/src/server"
	"github.com/labstack/echo/v4"
)

func HomeHandler(s server.Server) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(200, model.HomeResponse{
			Message: "Welcome to server in go",
			Status:  true,
		})
	}
}
