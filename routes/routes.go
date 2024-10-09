package routes

import (
	"belajar-go-fiber/controllers"
	"belajar-go-fiber/middleware"
	"github.com/gofiber/fiber/v2"
)

// RouteInit initializes the routes
func RouteInit(r *fiber.App, userController *controllers.UserController, controller *controllers.ProductController) {
	// Public routes
	r.Post("/register", userController.RegisterUser)
	r.Post("/login", userController.LoginUser)

	// Protected routes - Admin only
	protected := r.Group("/users", middleware.AuthMiddleware)

	protected.Get("/", middleware.AdminMiddleware, userController.GetAllUsers)
	protected.Get("/:id", middleware.AdminMiddleware, userController.GetUserByID)
	protected.Patch("/:id", middleware.AdminMiddleware, userController.UpdateUser)
	protected.Delete("/:id", middleware.AdminMiddleware, userController.DeleteUser)

	// Product routes
	r.Post("/products", middleware.AuthMiddleware, middleware.AdminMiddleware, controller.CreateProduct)
	r.Get("/products", controller.GetAllProducts)
}
