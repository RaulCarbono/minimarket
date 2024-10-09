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

func GetUserByIdHandler(s server.Server) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			fmt.Println("[error]:", err)
		}
		user, err := repository.GetUserById(ctx, id)
		if err != nil {
			http.Error(ctx.Response().Writer, err.Error(), http.StatusBadRequest)
			return err
		}
		userByIdResponse := &model.UserByIdResponse{
			Id:       int(user.ID),
			Email:    user.Email,
			Password: user.Password,
			Role:     user.Role,
		}
		return ctx.JSON(http.StatusOK, userByIdResponse)
	}
}
