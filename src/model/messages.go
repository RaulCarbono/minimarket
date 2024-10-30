package model

import (
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
)

type (
	HomeResponse struct {
		Message string `json:"message"`
		Status  bool   `json:"status"`
	}

	UserByIdResponse struct {
		Id       int    `json:"id"`
		Email    string `json:"email"`
		Password string `json:"password"`
		Role     string `json:"role"`
	}

	UsersResponse struct {
		Id       uint   `json:"id"`
		Email    string `json:"email"`
		Password string `json:"password"`
		Role     string `json:"role"`
	}

	UpdateResponse struct {
		Message string `json:"message"`
	}

	CustomerByIdResponse struct {
		Id       int      `json:"id"`
		Name     string   `json:"name"`
		LastName string   `json:"lastname"`
		Phone    string   `json:"phone"`
		Email    string   `json:"email"`
		Role     string   `json:"role"`
		Orders   []*Order `json:"oders"`
	}

	ProductResponse struct {
		Name        string  `json:"name"`
		Description string  `json:"description"`
		Price       float64 `json:"price"`
		Image       string  `json:"image"`
	}

	ProductByIdResponse struct {
		Id       uint `json:"id"`
		OrderId  uint `json:"orderId"`
		Producto *ProductResponse
	}

	LoginResponse struct {
		Token string `json:"token"`
	}

	SignUpResponse struct {
		Name  string `json:"name"`
		Email string `json:"email"`
		Role  string `json:"role"`
	}

	OrderResponse struct {
		OrderId uint
		Message string
	}

	AddItemResponse struct {
		ItemId  uint
		Message string
	}

	SignUpRequest struct {
		Name     string `json:"name" validate:"required"`
		LastName string `json:"lastname" validate:"required"`
		Phone    string `json:"phone" validate:"required"`
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
		Role     string `json:"role" validate:"required"`
	}

	LoginRequest struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
		Role     string `json:"role" validate:"required"`
	}

	ProductRequest struct {
		Name        string  `json:"name" validate:"required"`
		Description string  `json:"description" validate:"required"`
		Price       float64 `json:"price" validate:"required"`
		Image       string  `json:"image" validate:"required"`
	}

	UpdateProductRequest struct {
		Name        string  `json:"name"`
		Description string  `json:"description"`
		Price       float64 `json:"price"`
		Image       string  `json:"image"`
	}

	OrderRequest struct {
		CustomerId uint `json:"customerId" validate:"required"`
	}

	AddItemRequest struct {
		OrderId   uint `json:"orderId" validate:"required"`
		ProductId uint `json:"productId" validate:"required"`
		Amount    uint `json:"Amount" validate:"required"`
	}

	CustomValidator struct {
		Validator *validator.Validate
	}
)

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.Validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}
