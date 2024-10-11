package model

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	Name     string `gorm:"type:varchar(255);not null"`
	LastName string `gorm:"type:varchar(255);not null"`
	Phone    string `gorm:"type:varchar(255);not null"`
	User     User   `gorm:"foreignKey:ID"`
}
