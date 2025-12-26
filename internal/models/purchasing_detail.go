package models

type PurchasingDetail struct {
	ID           uint    `gorm:"primaryKey"`
	PurchasingID uint    `gorm:"not null;index"`
	ItemID       uint    `gorm:"not null"`
	Qty          int     `gorm:"not null"`
	Price        float64 `gorm:"not null"`
	SubTotal     float64 `gorm:"not null"`

	Item Item `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
}
