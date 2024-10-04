package repository

import (
	"context"

	"github.com/go/mini_market/src/model"
)

type Repository interface {
	GetUserById(ctx context.Context, id int) (*model.User, error)
	Close() error
}

var implementation Repository

func SetRepository(repository Repository) {
	implementation = repository
}

func GetUserById(ctx context.Context, id int) (*model.User, error) {
	return implementation.GetUserById(ctx, id)
}

func Close() error {
	return implementation.Close()
}
