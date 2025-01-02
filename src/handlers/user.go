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

func GetUserHandler(s server.Server) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		rows, err := repository.GetUsers(ctx)
		if err != nil {
			http.Error(ctx.Response().Writer, err.Error(), http.StatusBadRequest)
			return err
		}
		var usersResponse []*model.UsersResponse

		for _, user := range rows {
			usersResponse = append(usersResponse, &model.UsersResponse{
				Id:       user.ID,
				Email:    user.Email,
				Password: user.Password,
				Role:     user.Role,
			})

		}
		return ctx.JSON(http.StatusOK, usersResponse)
	}
}

func UpdateUserHandler(s server.Server) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		userId, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			return &echo.HTTPError{
				Code:    echo.ErrBadGateway.Code,
				Message: err,
			}
		}

		changes := new(model.User)
		if err := ctx.Bind(changes); err != nil {
			return &echo.HTTPError{
				Code:    echo.ErrBadRequest.Code,
				Message: err.Error(),
			}
		}
		err = repository.UpdateUser(ctx, userId, changes)

		if err != nil {
			return &echo.HTTPError{
				Code:    echo.ErrBadGateway.Code,
				Message: err,
			}
		}

		return ctx.JSON(http.StatusOK, &model.UpdateResponse{
			Message: "user successfully updated",
		})
	}
}

func DeleteUserHandler(s server.Server) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		userId, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			return &echo.HTTPError{
				Code:    echo.ErrBadGateway.Code,
				Message: err,
			}
		}

		enabled := new(model.User)

		if err := ctx.Validate(enabled); err != nil {
			return &echo.HTTPError{
				Code:    echo.ErrBadRequest.Code,
				Message: err.Error(),
			}
		}

		if err := ctx.Bind(enabled); err != nil {
			return &echo.HTTPError{
				Code:    echo.ErrBadRequest.Code,
				Message: err.Error(),
			}
		}

		err = repository.DeleteUser(ctx, userId)

		if err != nil {
			return &echo.HTTPError{
				Code:    echo.ErrBadGateway.Code,
				Message: err,
			}
		}

		return ctx.JSON(http.StatusOK, &model.UpdateResponse{
			Message: "user successfully Delete",
		})
	}
}
