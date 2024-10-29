package main

import (
	"belajar-go-fiber/config"
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
	config.InitRedis()
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
	articleRepo := repositories.NewArticleRepository(database.DB)
	articleService := services.NewArticleService(articleRepo)
	articleController := controllers.NewArticleController(articleService)
	// Initialize 		routes
	routes.RouteInit(app, userController, productController, articleController)

	// Start the server
	app.Listen(":3000")
}
