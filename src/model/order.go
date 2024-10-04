package model

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	CustomerId int `gorm:"not null"`
	OrderId    int `gorm:"not null"`
	ProductId  int `gorm:"not null"`
	Amount     int `gorm:"not null"`
}
