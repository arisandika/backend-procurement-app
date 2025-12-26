package handlers

import (
	"procurement-app/internal/services"
	"procurement-app/internal/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type PurchasingHandler struct {
	service *services.PurchasingService
}

func NewPurchasingHandler(service *services.PurchasingService) *PurchasingHandler {
	return &PurchasingHandler{service}
}

// CREATE (TRANSACTIONAL)
func (h *PurchasingHandler) Create(c *fiber.Ctx) error {
	var req services.PurchasingRequest

	if err := c.BodyParser(&req); err != nil {
		return utils.Error(
			c,
			400,
			"Invalid request body",
			"body",
			err.Error(),
		)
	}

	if req.SupplierID == 0 || req.UserID == 0 {
		return utils.Error(
			c,
			422,
			"Validation error",
			"supplier_id/user_id",
			"supplier_id and user_id are required",
		)
	}

	if len(req.Details) == 0 {
		return utils.Error(
			c,
			422,
			"Validation error",
			"details",
			"details cannot be empty",
		)
	}

	purchasing, err := h.service.Create(&req)
	if err != nil {
		return utils.Error(
			c,
			400,
			"Failed to create purchasing",
			"transaction",
			err.Error(),
		)
	}

	return utils.Success(
		c,
		201,
		"Purchasing created successfully",
		purchasing,
	)
}

// GET ALL
func (h *PurchasingHandler) GetAll(c *fiber.Ctx) error {
	data, err := h.service.GetAll()
	if err != nil {
		return utils.Error(
			c,
			500,
			"Failed to retrieve purchasings",
			"database",
			err.Error(),
		)
	}

	return utils.Success(
		c,
		200,
		"Purchasings retrieved successfully",
		data,
	)
}

// GET BY ID
func (h *PurchasingHandler) GetByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.Error(
			c,
			400,
			"Invalid purchasing ID",
			"id",
			"ID must be numeric",
		)
	}

	data, err := h.service.GetByID(uint(id))
	if err != nil {
		return utils.Error(
			c,
			404,
			"Purchasing not found",
			"id",
			err.Error(),
		)
	}

	return utils.Success(
		c,
		200,
		"Purchasing retrieved successfully",
		data,
	)
}
