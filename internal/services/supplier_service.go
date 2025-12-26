package services

import (
	"procurement-app/internal/models"
	"procurement-app/internal/repositories"

)

type SupplierService struct {
	repo *repositories.SupplierRepository
}

func NewSupplierService(repo *repositories.SupplierRepository) *SupplierService {
	return &SupplierService{repo}
}

// CREATE
func (s *SupplierService) Create(data *models.Supplier) error {
	return s.repo.Create(data)
}

// GET ALL
func (s *SupplierService) GetAll() ([]models.Supplier, error) {
	return s.repo.FindAll()
}

// GET BY ID
func (s *SupplierService) GetByID(id uint) (*models.Supplier, error) {
	return s.repo.FindByID(id)
}

// UPDATE
func (s *SupplierService) Update(id uint, data *models.Supplier) error {
	supplier, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	supplier.Name = data.Name
	supplier.Email = data.Email
	supplier.Address = data.Address

	return s.repo.Update(supplier)
}

// DELETE
func (s *SupplierService) Delete(id uint) error {
	return s.repo.Delete(id)
}
