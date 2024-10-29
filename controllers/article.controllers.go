package controllers

import (
	"belajar-go-fiber/models"
	"belajar-go-fiber/services"
	"belajar-go-fiber/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"log"
	"strconv"
)

type ArticleController struct {
	Service   services.ArticleService
	Validator *validator.Validate
}

func NewArticleController(service services.ArticleService) *ArticleController {
	return &ArticleController{
		Service:   service,
		Validator: validator.New(),
	}
}

// CreateArticle creates a new article
func (ac *ArticleController) CreateArticle(c *fiber.Ctx) error {
	article := new(models.Article)

	if err := c.BodyParser(article); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid input",
			"error":   err.Error(),
		})
	}

	if err := ac.Validator.Struct(article); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Validation failed",
			"error":   err.Error(),
		})
	}

	if err := ac.Service.CreateArticle(article); err != nil {
		log.Printf("CreateProduct service error: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(utils.Response{
			Status:  "error",
			Message: "Failed to create article",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(utils.Response{
		Status:  "success",
		Message: "Product created successfully",
		Data:    article,
	})
}

func (ac *ArticleController) GetAllArticles(c *fiber.Ctx) error {
	// Pagination parameters
	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil || page < 1 {
		page = 1 // Default to page 1 if parsing fails or invalid
	}

	limit, err := strconv.Atoi(c.Query("limit", "10"))
	if err != nil || limit < 1 {
		limit = 10 // Default to limit of 10 if parsing fails or invalid
	}

	// Get total count of products for pagination
	totalCount, err := ac.Service.Count() // You need to implement this method in your service
	if err != nil {
		log.Printf("GetAllArticles total count error: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(utils.Response{
			Status:  "error",
			Message: "Failed to fetch total articles count",
			Error:   err.Error(),
		})
	}

	// Fetch products with pagination
	articles, err := ac.Service.FindAll(page, limit)
	if err != nil {
		log.Printf("GetAllProducts service error: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(utils.Response{
			Status:  "error",
			Message: "Failed to fetch products",
			Error:   err.Error(),
		})
	}

	// Calculate total pages
	totalPages := (totalCount + limit - 1) / limit // Ceiling division

	return c.Status(fiber.StatusOK).JSON(utils.Response{
		Status:  "success",
		Message: "Articles retrieved successfully",
		Data:    articles,
		Meta: &utils.PaginationMeta{
			CurrentPage: page,
			TotalPages:  totalPages,
			TotalItems:  totalCount, // Add total items for completeness
		},
	})
}

func (ac *ArticleController) GetArticleByID(c *fiber.Ctx) error {
	id := c.Params("id")
	article, err := ac.Service.FindByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(utils.Response{
			Status:  "error",
			Message: "Article not found",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(utils.Response{
		Status: "success",
		Data:   article,
	})
}

func (ac *ArticleController) UpdateArticle(c *fiber.Ctx) error {

	idStr := c.Params("id")
	article := new(models.Article)
	// Parse the request body into the user struct
	if err := c.BodyParser(article); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.Response{
			Status:  "error",
			Message: "Invalid input",
		})
	}

	// Convert string ID to uuid
	id, err := uuid.Parse(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.Response{
			Status:  "error",
			Message: "Invalid UUID",
		})
	}
	article.ID = id // Set the ID for updating

	// Validate the user struct
	if err := ac.Validator.Struct(article); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.Response{
			Status:  "error",
			Message: utils.ParseValidationErrors(err.(validator.ValidationErrors)),
		})
	}

	// Call the service to update the user
	if err := ac.Service.UpdateArticle(article); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(utils.Response{
			Status:  "error",
			Message: "User not found or update failed",
		})
	}

	return c.Status(fiber.StatusOK).JSON(utils.Response{
		Status:  "success",
		Message: "User updated successfully",
	})

}

func (ac *ArticleController) DeleteArticle(c *fiber.Ctx) error {
	idStr := c.Params("id")
	err := ac.Service.DeleteArticle(idStr)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(utils.Response{
			Status:  "error",
			Message: "Article not found",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(utils.Response{
		Status:  "success",
		Message: "Article deleted successfully",
	})
}
