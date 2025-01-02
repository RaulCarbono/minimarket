package model

import (
	"gorm.io/gorm"
)

type (
	Product struct {
		gorm.Model
		Name        string  `gorm:"type:varchar(255);not null" json:"name"`
		Description string  `gorm:"type:varchar(255);not null" json:"description"`
		Price       float64 `gorm:"type:float;not null" json:"price"`
		Image       string  `gorm:"type:varchar(255);not null" json:"image"`
		Order       []Order `gorm:"many2many:OrderProduct;" json:"-"`
	}
)
