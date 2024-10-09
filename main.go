package main

import (
	"belajar-go-fiber/controllers"
	"belajar-go-fiber/database"
	"belajar-go-fiber/database/migrations"
	"belajar-go-fiber/repositories"
	"belajar-go-fiber/services"

	//"log"
	"belajar-go-fiber/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Initialize database
	database.DatabaseInit()
	// Inisialisasi Migration
	migrations.Migration()
	// Initialize repository and service
	userRepo := repositories.NewUserRepository(database.DB)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)
	productRepo := repositories.NewProductRepository(database.DB)
	productService := services.NewProductService(productRepo)
	productController := controllers.NewProductController(productService)
	// Initialize routes
	routes.RouteInit(app, userController, productController)

	// Start the server
	app.Listen(":3000")
}
