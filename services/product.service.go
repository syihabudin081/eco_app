package services

import (
	"belajar-go-fiber/config"
	"belajar-go-fiber/models"
	"belajar-go-fiber/repositories"
	"encoding/json"
	"errors"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2/log"
	"strconv"
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
	return p.repo.CreateProduct(product)
}

// GetAllProducts retrieves all products with pagination
func (p *productService) GetAllProducts(page, limit int) ([]*models.Product, error) {
	cacheKey := "products:page:" + strconv.Itoa(page) + ":limit:" + strconv.Itoa(limit)

	// Cek cache Redis
	val, err := config.RedisClient.Get(config.Ctx, cacheKey).Result()
	if err == redis.Nil {
		// Data tidak ada di cache, ambil dari PostgreSQL
		offset := (page - 1) * limit
		products, err := p.repo.FindAll(page, limit, offset) // Pass pagination parameters
		if err != nil {
			return nil, err
		}

		// Simpan hasil ke cache
		jsonData, marshalErr := json.Marshal(products)
		if marshalErr != nil {
			return nil, marshalErr
		}
		config.RedisClient.Set(config.Ctx, cacheKey, jsonData, 0) // 0 berarti tidak ada expiry

		return products, nil
	} else if err != nil {
		return nil, err
	}

	// Deserialize data dari Redis ke slice of Product
	log.Info("Deserialize data from redis")
	var products []*models.Product
	unmarshalErr := json.Unmarshal([]byte(val), &products)
	if unmarshalErr != nil {
		return nil, unmarshalErr
	}

	return products, nil
}

// GetTotalProductsCount retrieves the total number of products
func (p *productService) GetTotalProductsCount() (int, error) {
	return p.repo.Count()
}

// DeleteProduct deletes a product by ID
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
