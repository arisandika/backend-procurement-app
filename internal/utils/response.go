package utils

import "github.com/gofiber/fiber/v2"

type ErrorDetail struct {
	Field  string `json:"field"`
	Reason string `json:"reason"`
}

func Success(c *fiber.Ctx, code int, message string, data interface{}) error {
	return c.Status(code).JSON(fiber.Map{
		"code":    code,
		"message": message,
		"data":    data,
	})
}

func Error(c *fiber.Ctx, code int, message, field, reason string) error {
	return c.Status(code).JSON(fiber.Map{
		"code":    code,
		"message": message,
		"error": ErrorDetail{
			Field:  field,
			Reason: reason,
		},
	})
}
