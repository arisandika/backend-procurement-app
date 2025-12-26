package repositories

import (
	"procurement-app/internal/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ItemRepository struct {
	db *gorm.DB
}

func NewItemRepository(db *gorm.DB) *ItemRepository {
	return &ItemRepository{db}
}

// CREATE
func (r *ItemRepository) Create(item *models.Item) error {
	return r.db.Create(item).Error
}

// GET ALL
func (r *ItemRepository) FindAll() ([]models.Item, error) {
	var items []models.Item
	err := r.db.Find(&items).Error
	return items, err
}

// GET BY ID
func (r *ItemRepository) FindByID(id uint) (*models.Item, error) {
	var item models.Item
	if err := r.db.First(&item, id).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

// UPDATE
func (r *ItemRepository) Update(item *models.Item) error {
	return r.db.Save(item).Error
}

// DELETE
func (r *ItemRepository) Delete(id uint) error {
	return r.db.Delete(&models.Item{}, id).Error
}

// UPDATE STOCK
func (r *ItemRepository) WithTx(tx *gorm.DB) *ItemRepository {
	return &ItemRepository{db: tx}
}

func (r *ItemRepository) FindByIDForUpdate(id uint, tx *gorm.DB) (*models.Item, error) {
	var item models.Item
	if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&item, id).Error; err != nil {
		return nil, err
	}
	return &item, nil
}
