package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"size:50;uniqueIndex;not null"`
	Password string `gorm:"not null"`
	Role     string `gorm:"size:20;not null"`
}
