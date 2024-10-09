package services

import (
	"belajar-go-fiber/models"
	"belajar-go-fiber/repositories"
	"errors"
)

// ProductService interface defines methods for product-related business logic
type ProductService interface {
	CreateProduct(product *models.Product) error
	GetAllProducts(page, limit int) ([]*models.Product, error)
	GetProductByID(id string) (*models.Product, error)
	GetTotalProductsCount() (int, error)
	DeleteProduct(id string) error
}

// productService struct implements ProductService
type productService struct {
	repo repositories.ProductRepository
}

// NewProductService creates a new product service
func NewProductService(repo repositories.ProductRepository) ProductService {
	return &productService{repo: repo}
}

// CreateProduct handles product creation
func (p *productService) CreateProduct(product *models.Product) error {
	if product == nil {
		return errors.New("product cannot be nil")
	}
	// Perform business logic, validation, etc. before creating
	return p.repo.CreateProduct(product)
}

// GetAllProducts retrieves all products with pagination
func (p *productService) GetAllProducts(page, limit int) ([]*models.Product, error) {
	offset := (page - 1) * limit
	products, err := p.repo.FindAll(page, limit, offset) // Pass pagination parameters
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (p *productService) GetTotalProductsCount() (int, error) {
	return p.repo.Count()
}

func (p *productService) DeleteProduct(id string) error {
	return p.repo.Delete(id)
}

// GetProductByID retrieves a product by its ID
func (p *productService) GetProductByID(id string) (*models.Product, error) {
	product, err := p.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}
