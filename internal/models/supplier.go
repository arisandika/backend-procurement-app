package models

import "gorm.io/gorm"

type Supplier struct {
	gorm.Model
	Name    string `gorm:"not null"`
	Email   string `gorm:"not null"`
	Address string `gorm:"not null"`
}
