package handlers

import (
	"fmt"
	"net/http"

	"github.com/go/mini_market/src/model"
	"github.com/go/mini_market/src/repository"
	"github.com/go/mini_market/src/server"
	"github.com/labstack/echo/v4"
)

func InsertOrderHandler(s server.Server) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		order := new(model.OrderRequest)
		if err := ctx.Bind(order); err != nil {
			fmt.Println(err)
			return &echo.HTTPError{
				Code:    echo.ErrBadRequest.Code,
				Message: err,
			}
		}
		if err := ctx.Validate(order); err != nil {
			return &echo.HTTPError{
				Code:    echo.ErrBadRequest.Code,
				Message: err.Error(),
			}
		}

		var newOrder = &model.Order{
			CustomerId: order.CustomerId,
		}
		err := repository.InsertOrder(ctx, newOrder)
		if err != nil {
			return &echo.HTTPError{
				Code:    echo.ErrBadGateway.Code,
				Message: err,
			}
		}
		return ctx.JSON(http.StatusOK, &model.OrderResponse{
			OrderId: newOrder.ID,
			Message: "order successfully created",
		})
	}
}

func AddItemHandler(s server.Server) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		item := new(model.AddItemRequest)
		if err := ctx.Bind(item); err != nil {
			fmt.Println(err)
			return &echo.HTTPError{
				Code:    echo.ErrBadRequest.Code,
				Message: err,
			}
		}
		if err := ctx.Validate(item); err != nil {
			return &echo.HTTPError{
				Code:    echo.ErrBadRequest.Code,
				Message: err.Error(),
			}
		}

		var newItem = &model.OrderProduct{
			OrderId:   item.OrderId,
			ProductId: item.ProductId,
			Amount:    item.Amount,
		}
		err := repository.AddItem(ctx, newItem)
		if err != nil {
			return &echo.HTTPError{
				Code:    echo.ErrBadGateway.Code,
				Message: err,
			}
		}
		return ctx.JSON(http.StatusOK, &model.AddItemResponse{
			ItemId:  newItem.ID,
			Message: "order successfully created",
		})
	}
}
