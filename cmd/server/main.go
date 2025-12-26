package main

import (
	"log"
	"procurement-app/config"
	"procurement-app/internal/handlers"
	"procurement-app/internal/models"
	"procurement-app/internal/repositories"
	"procurement-app/internal/routes"
	"procurement-app/internal/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func main() {
	// LOAD .env
	if err := godotenv.Load(); err != nil {
		log.Fatal("Failed to load .env")
	}

	// CONNECT DB
	db := config.ConnectDB()

	// MIGRATE TABLES
	if err := db.AutoMigrate(
		&models.User{},
		&models.Supplier{},
		&models.Item{},
		&models.Purchasing{},
		&models.PurchasingDetail{},
	); err != nil {
		log.Fatalf("Failed to auto-migrate tables: %v", err)
	}
	log.Println("Database migration completed")

	// SEED DATA
	seedData(db)

	// REPOSITORIES
	userRepo := repositories.NewUserRepository(db)
	supplierRepo := repositories.NewSupplierRepository(db)
	itemRepo := repositories.NewItemRepository(db)
	purchasingRepo := repositories.NewPurchasingRepository(db)

	// SERVICES
	authService := services.NewAuthService(userRepo)
	supplierService := services.NewSupplierService(supplierRepo)
	itemService := services.NewItemService(itemRepo)
	purchasingService := services.NewPurchasingService(purchasingRepo)

	// HANDLERS
	authHandler := handlers.NewAuthHandler(authService)
	supplierHandler := handlers.NewSupplierHandler(supplierService)
	itemHandler := handlers.NewItemHandler(itemService)
	purchasingHandler := handlers.NewPurchasingHandler(purchasingService)

	// FIBER APP
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5500,http://127.0.0.1:5500",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
		AllowCredentials: true,
	}))

	// ROUTES
	routes.Setup(
		app,
		authHandler,
		supplierHandler,
		itemHandler,
		purchasingHandler,
	)

	log.Println("Server running on :3000")
	log.Fatal(app.Listen(":3000"))
}

// SEED DATA
func seedData(db *gorm.DB) {
	// ADMIN USER
	var count int64
	db.Model(&models.User{}).Count(&count)
	if count == 0 {
		pass, _ := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
		admin := models.User{
			Username: "admin",
			Password: string(pass),
			Role:     "admin",
		}
		db.Create(&admin)
		log.Println("Admin user created: username=admin, password=admin123")
	}

	// DEFAULT SUPPLIER
	db.FirstOrCreate(&models.Supplier{}, models.Supplier{
		Name:    "PT Abadi Jaya",
		Email:   "abadi@jaya.com",
		Address: "Jakarta Pusat",
	})

	// DEFAULT ITEM
	db.FirstOrCreate(&models.Item{}, models.Item{
		Name:  "MacBook Pro 16-inch",
		Price: 15000000,
		Stock: 10,
	})
	db.FirstOrCreate(&models.Item{}, models.Item{
		Name:  "DJI Mavic Air 2",
		Price: 5000000,
		Stock: 20,
	})
}
