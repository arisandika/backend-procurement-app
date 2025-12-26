package routes

import (
	"procurement-app/internal/handlers"
	"procurement-app/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func Setup(
	app *fiber.App,
	authHandler *handlers.AuthHandler,
	supplierHandler *handlers.SupplierHandler,
	itemHandler *handlers.ItemHandler,
	purchasingHandler *handlers.PurchasingHandler,
) {
	api := app.Group("/api")

	// PUBLIC
	api.Post("/register", authHandler.Register)
	api.Post("/login", authHandler.Login)

	api.Get("/profile", middleware.JWTProtected(), func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"user_id": c.Locals("user_id"),
			"role":    c.Locals("role"),
		})
	})

	// ADMIN ONLY
	admin := api.Group(
		"/admin",
		middleware.JWTProtected(),
		middleware.AdminOnly(),
	)

	admin.Get("/dashboard", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "welcome admin"})
	})

	// SUPPLIER CRUD
	admin.Post("/suppliers", supplierHandler.Create)
	admin.Get("/suppliers", supplierHandler.GetAll)
	admin.Get("/suppliers/:id", supplierHandler.GetByID)
	admin.Put("/suppliers/:id", supplierHandler.Update)
	admin.Delete("/suppliers/:id", supplierHandler.Delete)

	// ITEM CRUD
	admin.Post("/items", itemHandler.Create)
	admin.Get("/items", itemHandler.GetAll)
	admin.Get("/items/:id", itemHandler.GetByID)
	admin.Put("/items/:id", itemHandler.Update)
	admin.Delete("/items/:id", itemHandler.Delete)

	// PURCHASING CRUD
	admin.Post("/purchasings", purchasingHandler.Create)
	admin.Get("/purchasings", purchasingHandler.GetAll)
	admin.Get("/purchasings/:id", purchasingHandler.GetByID)
}
