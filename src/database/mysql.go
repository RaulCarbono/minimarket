package database

import (
	"context"
	"fmt"

	"github.com/go/mini_market/src/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MysqlRepositori struct {
	db *gorm.DB
}

func DBConnection(DSN string) (*MysqlRepositori, error) {
	db, err := gorm.Open(mysql.Open(DSN), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&model.User{}, &model.Product{}, &model.Order{}, &model.Customer{})
	return &MysqlRepositori{db: db}, nil
}

func (repo *MysqlRepositori) GetUserById(ctx context.Context, id int) (*model.User, error) {
	fmt.Print("User id")
	return nil, nil
}

func (repo *MysqlRepositori) Close() error {
	fmt.Print("error")
	return nil
}
