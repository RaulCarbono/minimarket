package repository

import (
	"github.com/go/mini_market/src/model"
	"github.com/labstack/echo/v4"
)

type Repository interface {
	GetUserById(ctx echo.Context, id int) (*model.User, error)
	GetUsers(ctx echo.Context) ([]*model.User, error)
	GetUserByEmail(ctx echo.Context, email string) (*model.User, error)
	InsertUser(ctx echo.Context, newUser *model.User) error
	UpdateUser(ctx echo.Context, userId int, changes interface{}) error
	DeleteUser(ctx echo.Context, userId int) error
	InsertCustomer(ctx echo.Context, newCustomer *model.Customer) error
	GetCustomerById(ctx echo.Context, id int) (*model.Customer, error)
	GetProductById(ctx echo.Context, id int) (*model.Product, error)
	GetProduct(ctx echo.Context) ([]*model.Product, error)
	InsertProduct(ctx echo.Context, newProduct *model.Product) error
	InsertOrder(ctx echo.Context, newOrder *model.Order) error
	AddItem(ctx echo.Context, newItem *model.OrderProduct) error
	Close() error
}

var implementation Repository

func SetRepository(repository Repository) {
	implementation = repository
}

func GetUserById(ctx echo.Context, id int) (*model.User, error) {
	return implementation.GetUserById(ctx, id)
}

func GetUsers(ctx echo.Context) ([]*model.User, error) {
	return implementation.GetUsers(ctx)
}

func GetUserByEmail(ctx echo.Context, email string) (*model.User, error) {
	return implementation.GetUserByEmail(ctx, email)
}

func InsertUser(ctx echo.Context, newUser *model.User) error {
	return implementation.InsertUser(ctx, newUser)
}

func UpdateUser(ctx echo.Context, userId int, changes interface{}) error {
	return implementation.UpdateUser(ctx, userId, changes)
}

func DeleteUser(ctx echo.Context, userId int) error {
	return implementation.DeleteUser(ctx, userId)
}

func GetCustomerById(ctx echo.Context, id int) (*model.Customer, error) {
	return implementation.GetCustomerById(ctx, id)
}

func InsertCustomer(ctx echo.Context, newCustomer *model.Customer) error {
	return implementation.InsertCustomer(ctx, newCustomer)
}

func GetProductById(ctx echo.Context, id int) (*model.Product, error) {
	return implementation.GetProductById(ctx, id)
}

func GetProduct(ctx echo.Context) ([]*model.Product, error) {
	return implementation.GetProduct(ctx)
}

func InsertProduct(ctx echo.Context, newProduct *model.Product) error {
	return implementation.InsertProduct(ctx, newProduct)
}

func InsertOrder(ctx echo.Context, newOrder *model.Order) error {
	return implementation.InsertOrder(ctx, newOrder)
}

func AddItem(ctx echo.Context, newItem *model.OrderProduct) error {
	return implementation.AddItem(ctx, newItem)
}

func Close() error {
	return implementation.Close()
}
