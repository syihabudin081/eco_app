package services

import (
	"belajar-go-fiber/models"
	"belajar-go-fiber/repositories"
	"fmt"
)

// UserService interface
type UserService interface {
	GetAllUsers(page, limit int) ([]*models.User, error)
	GetTotalUsersCount() (int, error)
	GetUserByID(id string) (*models.User, error)
	CreateUser(user *models.User) error
	UpdateUser(user *models.User) error
	DeleteUser(id string) error
	GetUserByEmail(email string) (*models.User, error)
}

// userService struct
type userService struct {
	repo repositories.UserRepository
}

// NewUserService creates a new user service
func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repo: repo}
}

// GetAllUsers retrieves all users with pagination
func (s *userService) GetAllUsers(page, limit int) ([]*models.User, error) {
	// Calculate the offset for pagination
	offset := (page - 1) * limit

	// Fetch users from repository with pagination
	return s.repo.FindAll(offset, limit)
}

func (s *userService) GetTotalUsersCount() (int, error) {
	return s.repo.Count()
}

// GetUserByID retrieves a user by ID
func (s *userService) GetUserByID(idStr string) (*models.User, error) {
	// Convert string ID to int

	user, err := s.repo.FindByID(idStr)
	if err != nil {
		return nil, fmt.Errorf("user not found: %v", err)
	}
	return user, nil
}

// CreateUser creates a new user in the database
func (s *userService) CreateUser(user *models.User) error {
	return s.repo.Create(user)
}

// GetUserByEmail fetches a user by email
func (s *userService) GetUserByEmail(email string) (*models.User, error) {
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return nil, fmt.Errorf("user not found by email: %v", err)
	}
	return user, nil
}

// UpdateUser updates an existing user
func (s *userService) UpdateUser(user *models.User) error {
	return s.repo.Update(user)
}

// DeleteUser deletes a user by ID
func (s *userService) DeleteUser(idStr string) error {

	return s.repo.Delete(idStr)
}
