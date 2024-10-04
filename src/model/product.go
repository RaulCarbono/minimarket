package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string  `gorm:"type:varchar(255);not null"`
	Description string  `gorm:"type:varchar(255);not null"`
	Price       float64 `gorm:"type:float;not null"`
	Image       string  `gorm:"type:varchar(255);not null"`
}
