package controllers

import (
	"belajar-go-fiber/models"
	"belajar-go-fiber/services"
	"belajar-go-fiber/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"log"
	"strconv"
)

// ProductController struct
type ProductController struct {
	Service   services.ProductService
	Validator *validator.Validate
}

// NewProductController creates a new ProductController
func NewProductController(service services.ProductService) *ProductController {
	return &ProductController{
		Service:   service,
		Validator: validator.New(),
	}
}

// CreateProduct creates a new product
func (pc *ProductController) CreateProduct(c *fiber.Ctx) error {
	product := new(models.Product)

	// Parse the request body into the product struct
	if err := c.BodyParser(product); err != nil {
		log.Printf("BodyParser error: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(utils.Response{
			Status:  "error",
			Message: "Invalid input",
			Error:   err.Error(),
		})
	}

	// Validate the product struct
	if err := pc.Validator.Struct(product); err != nil {
		validationErrors := utils.ParseValidationErrors(err)
		return c.Status(fiber.StatusBadRequest).JSON(utils.Response{
			Status:  "error",
			Message: "Validation failed",
			Error:   validationErrors,
		})
	}

	// Create the product
	if err := pc.Service.CreateProduct(product); err != nil {
		log.Printf("CreateProduct service error: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(utils.Response{
			Status:  "error",
			Message: "Failed to create product",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(utils.Response{
		Status:  "success",
		Message: "Product created successfully",
		Data:    product,
	})
}

// GetAllProducts retrieves all products with optional pagination
func (pc *ProductController) GetAllProducts(c *fiber.Ctx) error {
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
	totalCount, err := pc.Service.GetTotalProductsCount() // You need to implement this method in your service
	if err != nil {
		log.Printf("GetAllProducts total count error: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(utils.Response{
			Status:  "error",
			Message: "Failed to fetch total products count",
			Error:   err.Error(),
		})
	}

	// Fetch products with pagination
	products, err := pc.Service.GetAllProducts(page, limit)
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
		Message: "Products retrieved successfully",
		Data:    products,
		Meta: &utils.PaginationMeta{
			CurrentPage: page,
			TotalPages:  totalPages,
			TotalItems:  totalCount, // Add total items for completeness
		},
	})
}

func (pc *ProductController) GetProductByID(c *fiber.Ctx) error {
	id := c.Params("id")

	product, err := pc.Service.GetProductByID(id)
	if err != nil {
		log.Printf("GetProductByID service error: %v", err)
		return c.Status(fiber.StatusNotFound).JSON(utils.Response{
			Status:  "error",
			Message: "Product not found",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(utils.Response{
		Status:  "success",
		Message: "Product retrieved successfully",
		Data:    product,
	})
}

func (pc *ProductController) DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")

	err := pc.Service.DeleteProduct(id)
	if err != nil {
		log.Printf("DeleteProduct service error: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(utils.Response{
			Status:  "error",
			Message: "Failed to delete product",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(utils.Response{
		Status:  "success",
		Message: "Product deleted successfully",
	})
}
