package models

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	Name  string  `gorm:"not null"`
	Stock int     `gorm:"not null"`
	Price float64 `gorm:"not null"`
}
