package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/go/mini_market/src/model"
	"github.com/go/mini_market/src/repository"
	"github.com/go/mini_market/src/server"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

const (
	HASH_COST = 8
)

var (
	ErrValidateStruct = errors.New("some fields are required")
)

func SignUpHandler(s server.Server) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var validate = validator.New()
		var request = model.SignUpLoginRequest{}
		err := json.NewDecoder(ctx.Request().Body).Decode(&request)

		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid request payload")
		}

		if err := validate.Struct(request); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, ErrValidateStruct)
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), HASH_COST)

		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to hash password")
		}

		var newUser = model.User{
			Email:    request.Email,
			Password: string(hashedPassword),
			Role:     request.Role,
		}

		err = repository.InsertUser(ctx, &newUser)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create user")
		}
		return ctx.JSON(http.StatusOK, model.SignUpResponse{
			Email: newUser.Email,
			Role:  newUser.Role,
		})
	}
}
