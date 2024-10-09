package repositories

import (
	"belajar-go-fiber/models"
	"errors"
	"gorm.io/gorm"
)

// ProductRepository interface defines methods for product data access
type ProductRepository interface {
	CreateProduct(product *models.Product) error
	FindAll(page, limit, offset int) ([]*models.Product, error) // Add pagination parameters
	FindByID(id string) (*models.Product, error)
	Count() (int, error)
	Delete(id string) error
}

// productRepository struct implements ProductRepository
type productRepository struct {
	DB *gorm.DB
}

// NewProductRepository creates a new product repository
func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{DB: db}
}

// CreateProduct creates a new product
func (pr *productRepository) CreateProduct(product *models.Product) error {
	// Use a transaction to ensure atomicity
	return pr.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(product).Error; err != nil {
			return err
		}
		return nil
	})
}

// FindAll retrieves all products with pagination
func (pr *productRepository) FindAll(page, limit, offset int) ([]*models.Product, error) {
	var products []*models.Product
	err := pr.DB.Offset(offset).Limit(limit).Find(&products).Error // Use Offset and Limit for pagination
	return products, err
}

func (pr *productRepository) Count() (int, error) {
	var count int64
	err := pr.DB.Model(&models.Product{}).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return int(count), nil
}

// FindByID retrieves a product by its ID
func (pr *productRepository) FindByID(id string) (*models.Product, error) {
	var product models.Product
	err := pr.DB.First(&product, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("product not found")
	} else if err != nil {
		return nil, err
	}
	return &product, nil
}

func (pr *productRepository) Delete(id string) error {
	return pr.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("id = ?", id).Delete(&models.Product{}).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return errors.New("product not found")
			}
			return err
		}
		return nil
	})
}
