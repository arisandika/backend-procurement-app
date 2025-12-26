package handlers

import (
	"procurement-app/internal/models"
	"procurement-app/internal/services"
	"procurement-app/internal/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type ItemHandler struct {
	service *services.ItemService
}

func NewItemHandler(service *services.ItemService) *ItemHandler {
	return &ItemHandler{service}
}

// CREATE
func (h *ItemHandler) Create(c *fiber.Ctx) error {
	var item models.Item

	if err := c.BodyParser(&item); err != nil {
		return utils.Error(
			c,
			400,
			"Invalid request body",
			"body",
			err.Error(),
		)
	}

	requiredFields := map[string]interface{}{
		"name":  item.Name,
		"stock": item.Stock,
		"price": item.Price,
	}

	for field, value := range requiredFields {
		switch v := value.(type) {
		case string:
			if v == "" {
				return utils.Error(c, 422, "Validation error", field, field+" is required")
			}
		case int:
			if v < 0 {
				return utils.Error(c, 422, "Validation error", field, field+" must be >= 0")
			}
		case float64:
			if v <= 0 {
				return utils.Error(c, 422, "Validation error", field, field+" must be > 0")
			}
		}
	}

	if err := h.service.Create(&item); err != nil {
		return utils.Error(
			c,
			400,
			"Failed to create item",
			"database",
			err.Error(),
		)
	}

	return utils.Success(
		c,
		201,
		"Item created successfully",
		item,
	)
}

// GET ALL
func (h *ItemHandler) GetAll(c *fiber.Ctx) error {
	items, err := h.service.GetAll()
	if err != nil {
		return utils.Error(
			c,
			500,
			"Failed to retrieve items",
			"database",
			err.Error(),
		)
	}

	return utils.Success(
		c,
		200,
		"Items retrieved successfully",
		items,
	)
}

// GET BY ID
func (h *ItemHandler) GetByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.Error(
			c,
			400,
			"Invalid item ID",
			"id",
			"ID must be a number",
		)
	}

	item, err := h.service.GetByID(uint(id))
	if err != nil {
		return utils.Error(
			c,
			404,
			"Item not found",
			"id",
			err.Error(),
		)
	}

	return utils.Success(
		c,
		200,
		"Item retrieved successfully",
		item,
	)
}

// UPDATE
func (h *ItemHandler) Update(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.Error(
			c,
			400,
			"Invalid item ID",
			"id",
			"ID must be a number",
		)
	}

	var item models.Item
	if err := c.BodyParser(&item); err != nil {
		return utils.Error(
			c,
			400,
			"Invalid request body",
			"body",
			err.Error(),
		)
	}

	if err := h.service.Update(uint(id), &item); err != nil {
		return utils.Error(
			c,
			400,
			"Failed to update item",
			"database",
			err.Error(),
		)
	}

	return utils.Success(
		c,
		200,
		"Item updated successfully",
		item,
	)
}

// DELETE
func (h *ItemHandler) Delete(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.Error(
			c,
			400,
			"Invalid item ID",
			"id",
			"ID must be a number",
		)
	}

	if err := h.service.Delete(uint(id)); err != nil {
		return utils.Error(
			c,
			400,
			"Failed to delete item",
			"database",
			err.Error(),
		)
	}

	return utils.Success(
		c,
		200,
		"Item deleted successfully",
		nil,
	)
}
