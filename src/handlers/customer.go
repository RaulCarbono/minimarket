package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go/mini_market/src/model"
	"github.com/go/mini_market/src/repository"
	"github.com/go/mini_market/src/server"
	"github.com/labstack/echo/v4"
)

func GetCustomerByIdHandler(s server.Server) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			fmt.Println("[error]:", err)
		}
		echo.Logger.Info(echo.New().Logger, id)
		customer, err := repository.GetCustomerById(ctx, id)
		if err != nil {
			http.Error(ctx.Response().Writer, err.Error(), http.StatusBadRequest)
			return err
		}

		return ctx.JSON(http.StatusOK, &model.CustomerByIdResponse{
			Id:       int(customer.ID),
			Name:     customer.Name,
			LastName: customer.LastName,
			Phone:    customer.Phone,
			Email:    customer.User.Email,
			Role:     customer.User.Role,
			Orders:   customer.Orders,
		})
	}
}
