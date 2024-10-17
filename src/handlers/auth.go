package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/go/mini_market/src/model"
	"github.com/go/mini_market/src/repository"
	"github.com/go/mini_market/src/server"
	"github.com/golang-jwt/jwt"
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
		var request = model.SignUpRequest{}
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

		var newCustomer = model.Customer{
			Name:     request.Name,
			LastName: request.LastName,
			Phone:    request.Phone,
			User: model.User{
				Email:    request.Email,
				Password: string(hashedPassword),
				Role:     request.Role,
			},
		}

		err = repository.InsertCustomer(ctx, &newCustomer)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create user")
		}
		return ctx.JSON(http.StatusOK, model.SignUpResponse{
			Name:  newCustomer.Name,
			Email: newCustomer.User.Email,
		})
	}
}

func LoginHandler(s server.Server) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var validate = validator.New()
		var request = model.LoginRequest{}
		err := json.NewDecoder(ctx.Request().Body).Decode(&request)

		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid request payload")
		}

		if err := validate.Struct(request); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, ErrValidateStruct)
		}

		user, err := repository.GetUserByEmail(ctx, request.Email)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		if user == nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid credentials")
		}
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid credentials")
		}

		claims := model.AppClaims{
			UserId: fmt.Sprint(user.ID),
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(2 * time.Hour * 24).Unix(),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString([]byte(s.Config().JWTSecret))

		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		return ctx.JSON(http.StatusOK, model.LoginResponse{
			Token: tokenString,
		})
	}
}
