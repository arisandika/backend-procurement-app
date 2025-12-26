package models

import "time"

type Purchasing struct {
	ID         uint      `gorm:"primaryKey"`
	Date       time.Time `gorm:"not null"`
	SupplierID uint      `gorm:"not null"`
	UserID     uint      `gorm:"not null"`
	GrandTotal float64   `gorm:"not null"`

	Supplier Supplier           `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	User     User               `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	Details  []PurchasingDetail `gorm:"constraint:OnDelete:CASCADE;"`
}
