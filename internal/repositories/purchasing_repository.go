package repositories

import (
	"procurement-app/internal/models"

	"gorm.io/gorm"
)

type PurchasingRepository struct {
	db *gorm.DB
}

func NewPurchasingRepository(db *gorm.DB) *PurchasingRepository {
	return &PurchasingRepository{db}
}

// BEGIN TRANSACTION
func (r *PurchasingRepository) Begin() *gorm.DB {
	return r.db.Begin()
}

// INSERT HEADER
func (r *PurchasingRepository) CreatePurchasing(tx *gorm.DB, p *models.Purchasing) error {
	return tx.Create(p).Error
}

// INSERT DETAIL
func (r *PurchasingRepository) CreateDetail(tx *gorm.DB, d *models.PurchasingDetail) error {
	return tx.Create(d).Error
}

// UPDATE ITEM STOCK
func (r *PurchasingRepository) UpdateItemStock(tx *gorm.DB, itemID uint, qty int) error {
	return tx.Model(&models.Item{}).
		Where("id = ?", itemID).
		Update("stock", gorm.Expr("stock + ?", qty)).
		Error
}

// FIND ITEM
func (r *PurchasingRepository) FindItem(tx *gorm.DB, id uint) (*models.Item, error) {
	var item models.Item
	if err := tx.First(&item, id).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *PurchasingRepository) DB() *gorm.DB {
	return r.db
}
