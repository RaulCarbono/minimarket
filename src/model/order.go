package model

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	CustomerId uint      `json:"-"`
	Customer   *Customer `json:"-"`
	Product    []Product `gorm:"many2many:OrderProduct;"`
}
