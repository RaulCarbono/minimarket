package database

import (
	"github.com/go/mini_market/src/model"
	"github.com/labstack/echo/v4"
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

func (repo *MysqlRepositori) GetUserById(ctx echo.Context, id int) (*model.User, error) {
	var user model.User
	err := repo.db.Table("users").Select("id", "email", "password", "role").Where("id = ?", id).Scan(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *MysqlRepositori) GetUserByEmail(ctx echo.Context, email string) (*model.User, error) {
	var user model.User
	err := repo.db.Table("users").Select("id", "email", "password", "role").Where("email = ?", email).Scan(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *MysqlRepositori) InsertUser(ctx echo.Context, newUser *model.User) error {
	result := repo.db.Create(&newUser)
	return result.Error
}

func (repo *MysqlRepositori) Close() error {
	db, err := repo.db.DB()
	if err != nil {
		return err
	}
	return db.Close()
}
