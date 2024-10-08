package repository

import (
	"github.com/go/mini_market/src/model"
	"github.com/labstack/echo/v4"
)

type Repository interface {
	GetUserById(ctx echo.Context, id int) (*model.User, error)
	InsertUser(ctx echo.Context, newUser *model.User) error
	Close() error
}

var implementation Repository

func SetRepository(repository Repository) {
	implementation = repository
}

func GetUserById(ctx echo.Context, id int) (*model.User, error) {
	return implementation.GetUserById(ctx, id)
}

func InsertUser(ctx echo.Context, newUser *model.User) error {
	return implementation.InsertUser(ctx, newUser)
}

func Close() error {
	return implementation.Close()
}
