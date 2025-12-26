package handlers

import (
	"procurement-app/internal/models"
	"procurement-app/internal/services"
	"procurement-app/internal/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"

)

type SupplierHandler struct {
	service *services.SupplierService
}

func NewSupplierHandler(service *services.SupplierService) *SupplierHandler {
	return &SupplierHandler{service}
}

// CREATE
func (h *SupplierHandler) Create(c *fiber.Ctx) error {
	var supplier models.Supplier

	if err := c.BodyParser(&supplier); err != nil {
		return utils.Error(
			c,
			400,
			"Invalid request body",
			"body",
			err.Error(),
		)
	}

	// VALIDATE REQUIRED FIELDS
	requiredFields := map[string]string{
		"name":    supplier.Name,
		"email":   supplier.Email,
		"address": supplier.Address,
	}

	for field, value := range requiredFields {
		if value == "" {
			return utils.Error(
				c,
				422,
				"Validation error",
				field,
				field+" is required",
			)
		}
	}

	if err := h.service.Create(&supplier); err != nil {
		return utils.Error(
			c,
			400,
			"Failed to create supplier",
			"database",
			err.Error(),
		)
	}

	return utils.Success(
		c,
		201,
		"Supplier created successfully",
		supplier,
	)
}

// GET ALL
func (h *SupplierHandler) GetAll(c *fiber.Ctx) error {
	suppliers, err := h.service.GetAll()
	if err != nil {
		return utils.Error(
			c,
			500,
			"Failed to retrieve suppliers",
			"database",
			err.Error(),
		)
	}

	return utils.Success(
		c,
		200,
		"Suppliers retrieved successfully",
		suppliers,
	)
}

// GET BY ID
func (h *SupplierHandler) GetByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.Error(
			c,
			400,
			"Invalid supplier ID",
			"id",
			"ID must be a number",
		)
	}

	supplier, err := h.service.GetByID(uint(id))
	if err != nil {
		return utils.Error(
			c,
			404,
			"Supplier not found",
			"id",
			err.Error(),
		)
	}

	return utils.Success(
		c,
		200,
		"Supplier retrieved successfully",
		supplier,
	)
}

// UPDATE
func (h *SupplierHandler) Update(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.Error(
			c,
			400,
			"Invalid supplier ID",
			"id",
			"ID must be a number",
		)
	}

	var supplier models.Supplier

	if err := c.BodyParser(&supplier); err != nil {
		return utils.Error(
			c,
			400,
			"Invalid request body",
			"body",
			err.Error(),
		)
	}

	// VALIDATE REQUIRED FIELDS
	requiredFields := map[string]string{
		"name":    supplier.Name,
		"email":   supplier.Email,
		"address": supplier.Address,
	}

	for field, value := range requiredFields {
		if value == "" {
			return utils.Error(
				c,
				422,
				"Validation error",
				field,
				field+" is required",
			)
		}
	}

	if err := h.service.Update(uint(id), &supplier); err != nil {
		return utils.Error(
			c,
			400,
			"Failed to update supplier",
			"database",
			err.Error(),
		)
	}

	return utils.Success(
		c,
		200,
		"Supplier updated successfully",
		supplier,
	)
}

// DELETE
func (h *SupplierHandler) Delete(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.Error(
			c,
			400,
			"Invalid supplier ID",
			"id",
			"ID must be a number",
		)
	}

	if err := h.service.Delete(uint(id)); err != nil {
		return utils.Error(
			c,
			400,
			"Failed to delete supplier",
			"database",
			err.Error(),
		)
	}

	return utils.Success(
		c,
		200,
		"Supplier deleted successfully",
		nil,
	)
}
