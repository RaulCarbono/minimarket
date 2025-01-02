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

func GetProductByIdHandler(s server.Server) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			return echo.ErrBadGateway
		}

		product, err := repository.GetProductById(ctx, id)
		if err != nil {
			return echo.ErrNotFound
		}
		return ctx.JSON(http.StatusOK, &model.ProductByIdResponse{
			Id: product.ID,
			Producto: &model.ProductResponse{
				Name:        product.Name,
				Description: product.Description,
				Price:       product.Price,
				Image:       product.Image,
			},
		})
	}
}

func GetProductdHandler(s server.Server) echo.HandlerFunc {
	return func(ctx echo.Context) error {

		product, err := repository.GetProduct(ctx)
		if err != nil {
			return echo.ErrNotFound
		}
		return ctx.JSON(http.StatusOK, product)
	}
}

func InsertProductHandler(s server.Server) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		product := new(model.ProductRequest)
		if err := ctx.Bind(product); err != nil {
			fmt.Println(err)
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		if err := ctx.Validate(product); err != nil {
			echo.Logger.Error(echo.New().Logger, err.Error())
			return echo.NewHTTPError(http.StatusBadRequest, "some fields are required")
		}
		var newProduct = &model.Product{
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			Image:       product.Image,
		}
		err := repository.InsertProduct(ctx, newProduct)

		if err != nil {
			echo.Logger.Error(echo.New().Logger, err.Error())
			return echo.ErrBadGateway
		}
		return ctx.JSON(http.StatusOK, &model.ProductResponse{
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			Image:       product.Image,
		})
	}
}
