package routes

import (
	"belajar-go-fiber/controllers"
	"belajar-go-fiber/middleware"
	"github.com/gofiber/fiber/v2"
)

// RouteInit initializes the routes
func RouteInit(r *fiber.App, userController *controllers.UserController, controller *controllers.ProductController, articleController *controllers.ArticleController) {
	// Public routes
	r.Post("/register", userController.RegisterUser)
	r.Post("/login", userController.LoginUser)

	// Protected User routes - Admin only
	user := r.Group("/users", middleware.AuthMiddleware)
	user.Get("/", middleware.AdminMiddleware, userController.GetAllUsers)
	user.Get("/:id", middleware.AdminMiddleware, userController.GetUserByID)
	user.Patch("/:id", middleware.AdminMiddleware, userController.UpdateUser)
	user.Delete("/:id", middleware.AdminMiddleware, userController.DeleteUser)

	// Product routes
	product := r.Group("/product", middleware.AuthMiddleware)
	product.Post("/", middleware.AuthMiddleware, middleware.AdminMiddleware, controller.CreateProduct)
	product.Get("/", controller.GetAllProducts)
	product.Get("/:id", controller.GetProductByID)
	product.Delete("/:id", middleware.AdminMiddleware, controller.DeleteProduct)

	// Article routes
	article := r.Group("/article", middleware.AuthMiddleware)
	article.Post("/", middleware.AuthMiddleware, middleware.AdminMiddleware, articleController.CreateArticle)
	article.Get("/", articleController.GetAllArticles)
	article.Get("/:id", articleController.GetArticleByID)
	article.Patch("/:id", middleware.AdminMiddleware, articleController.UpdateArticle)
	article.Delete("/:id", middleware.AdminMiddleware, articleController.DeleteArticle)
}
