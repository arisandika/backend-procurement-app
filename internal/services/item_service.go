package services

import (
	"procurement-app/internal/models"
	"procurement-app/internal/repositories"
)

type ItemService struct {
	repo *repositories.ItemRepository
}

func NewItemService(repo *repositories.ItemRepository) *ItemService {
	return &ItemService{repo}
}

// CREATE
func (s *ItemService) Create(data *models.Item) error {
	return s.repo.Create(data)
}

// GET ALL
func (s *ItemService) GetAll() ([]models.Item, error) {
	return s.repo.FindAll()
}

// GET BY ID
func (s *ItemService) GetByID(id uint) (*models.Item, error) {
	return s.repo.FindByID(id)
}

// UPDATE
func (s *ItemService) Update(id uint, data *models.Item) error {
	item, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	item.Name = data.Name
	item.Stock = data.Stock
	item.Price = data.Price

	return s.repo.Update(item)
}

// DELETE
func (s *ItemService) Delete(id uint) error {
	return s.repo.Delete(id)
}
