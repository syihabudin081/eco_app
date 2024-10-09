package repositories

import (
	"belajar-go-fiber/models"
	"errors"
	"gorm.io/gorm"
)

// UserRepository interface defines the methods for user data access
type UserRepository interface {
	FindAll(offset, limit int) ([]*models.User, error)
	Count() (int, error)
	FindByID(id string) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	Create(user *models.User) error
	Update(user *models.User) error
	Delete(id string) error
}

// userRepository struct implements the UserRepository interface
type userRepository struct {
	DB *gorm.DB
}

// NewUserRepository creates a new instance of userRepository
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{DB: db}
}

// FindAll retrieves all users from the database with pagination
func (ur *userRepository) FindAll(offset, limit int) ([]*models.User, error) {
	var users []*models.User
	err := ur.DB.Preload("Role").Offset(offset).Limit(limit).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

// Count returns the total number of users
func (ur *userRepository) Count() (int, error) {
	var count int64
	err := ur.DB.Model(&models.User{}).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return int(count), nil
}

// FindByID retrieves a user by ID
func (ur *userRepository) FindByID(id string) (*models.User, error) {
	var user models.User
	err := ur.DB.Preload("Role").First(&user, "id = ?", id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("user not found")
	} else if err != nil {
		return nil, err
	}
	return &user, nil
}

// FindByEmail retrieves a user by their email
func (ur *userRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := ur.DB.Preload("Role").First(&user, "email = ?", email).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("user not found")
	} else if err != nil {
		return nil, err
	}
	return &user, nil
}

// Create adds a new user to the database
func (ur *userRepository) Create(user *models.User) error {
	return ur.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(user).Error; err != nil {
			return err
		}
		return nil
	})
}

// Update modifies an existing user in the database
func (ur *userRepository) Update(user *models.User) error {
	return ur.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(user).Error; err != nil {
			return err
		}
		return nil
	})
}

func (ur *userRepository) Delete(idStr string) error {
	// Use the Delete method, which automatically sets the DeletedAt field
	return ur.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("id = ?", idStr).Delete(&models.User{}).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return errors.New("user not found")
			}
			return err
		}
		return nil
	})
}
