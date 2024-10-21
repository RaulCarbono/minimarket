package model

import "gorm.io/gorm"

type OrderProduct struct {
	gorm.Model
	OrderId   uint
	ProductId uint
	Amount    uint
}
