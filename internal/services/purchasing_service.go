package services

import (
	"errors"
	"procurement-app/internal/dto"
	"procurement-app/internal/models"
	"procurement-app/internal/repositories"
	"time"
)

type PurchasingService struct {
	repo *repositories.PurchasingRepository
}

func NewPurchasingService(repo *repositories.PurchasingRepository) *PurchasingService {
	return &PurchasingService{repo}
}

// REQUEST DTO (payload dari client)
type PurchasingRequest struct {
	SupplierID uint `json:"supplier_id"`
	UserID     uint `json:"user_id"`
	Details    []struct {
		ItemID uint `json:"item_id"`
		Qty    int  `json:"qty"`
	} `json:"details"`
}

func (s *PurchasingService) Create(req *PurchasingRequest) (*models.Purchasing, error) {
	tx := s.repo.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var grandTotal float64

	purchasing := &models.Purchasing{
		Date:       time.Now(),
		SupplierID: req.SupplierID,
		UserID:     req.UserID,
	}

	// INSERT HEADER
	if err := s.repo.CreatePurchasing(tx, purchasing); err != nil {
		tx.Rollback()
		return nil, err
	}

	// LOOP DETAIL
	for _, d := range req.Details {

		if d.Qty <= 0 {
			tx.Rollback()
			return nil, errors.New("qty must be greater than zero")
		}

		item, err := s.repo.FindItem(tx, d.ItemID)
		if err != nil {
			tx.Rollback()
			return nil, errors.New("item not found")
		}

		subTotal := float64(d.Qty) * item.Price
		grandTotal += subTotal

		detail := models.PurchasingDetail{
			PurchasingID: purchasing.ID,
			ItemID:       item.ID,
			Qty:          d.Qty,
			Price:        item.Price,
			SubTotal:     subTotal,
		}

		// INSERT DETAIL
		if err := s.repo.CreateDetail(tx, &detail); err != nil {
			tx.Rollback()
			return nil, err
		}

		// UPDATE STOCK
		if err := s.repo.UpdateItemStock(tx, item.ID, d.Qty); err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	// UPDATE GRAND TOTAL
	if err := tx.Model(&models.Purchasing{}).
		Where("id = ?", purchasing.ID).
		Update("grand_total", grandTotal).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// COMMIT
	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	// SEND WEBHOOK
	go TriggerWebhook(purchasing.ID)

	return purchasing, nil
}

// GET BY ID
func (s *PurchasingService) GetByID(id uint) (dto.PurchasingResponse, error) {
	var p models.Purchasing
	err := s.repo.DB().
		Preload("Supplier").
		Preload("User").
		Preload("Details.Item").
		First(&p, id).Error
	if err != nil {
		return dto.PurchasingResponse{}, err
	}

	return mapToPurchasingResponse(p), nil
}

// GET ALL
func (s *PurchasingService) GetAll() ([]dto.PurchasingResponse, error) {
	var list []models.Purchasing
	err := s.repo.DB().
		Preload("Supplier").
		Preload("User").
		Preload("Details.Item").
		Find(&list).Error
	if err != nil {
		return nil, err
	}

	res := []dto.PurchasingResponse{}
	for _, p := range list {
		res = append(res, mapToPurchasingResponse(p))
	}

	return res, nil
}

// MAP TO DTO
func mapToPurchasingResponse(p models.Purchasing) dto.PurchasingResponse {
	details := []dto.PurchasingDetailResponse{}

	for _, d := range p.Details {
		details = append(details, dto.PurchasingDetailResponse{
			Qty:      d.Qty,
			Price:    d.Price,
			SubTotal: d.SubTotal,
			Item: dto.ItemResponse{
				ID:    d.Item.ID,
				Name:  d.Item.Name,
				Price: d.Item.Price,
			},
		})
	}

	return dto.PurchasingResponse{
		ID:         p.ID,
		Date:       p.Date,
		GrandTotal: p.GrandTotal,
		Supplier: dto.SupplierResponse{
			ID:      p.Supplier.ID,
			Name:    p.Supplier.Name,
			Address: p.Supplier.Address,
		},
		User: dto.UserResponse{
			ID:       p.User.ID,
			Username: p.User.Username,
		},
		Details: details,
	}
}
