package controllers

import (
	"belajar-go-fiber/models"
	"belajar-go-fiber/services"
	"belajar-go-fiber/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"log"
	"strconv"
)

// UserController struct
type UserController struct {
	Service   services.UserService
	Validator *validator.Validate
}

// NewUserController creates a new UserController
func NewUserController(service services.UserService) *UserController {
	return &UserController{
		Service:   service,
		Validator: validator.New(),
	}
}

// GetAllUsers retrieves all users with optional pagination
func (uc *UserController) GetAllUsers(c *fiber.Ctx) error {
	// Pagination parameters
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))

	// Get total count of users for pagination
	totalCount, err := uc.Service.GetTotalUsersCount()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.Response{
			Status:  "error",
			Message: "Error fetching total users count",
			Error:   err.Error(),
		})
	}

	// Fetch users with pagination
	users, err := uc.Service.GetAllUsers(page, limit)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.Response{
			Status:  "error",
			Message: "Error fetching users",
			Error:   err.Error(),
		})
	}

	// Calculate total pages
	totalPages := (totalCount + limit - 1) / limit // Ceiling division

	return c.Status(fiber.StatusOK).JSON(utils.Response{
		Status:  "success",
		Message: "Success Get All Users",
		Data:    users,
		Meta: &utils.PaginationMeta{ // Set the pagination metadata
			CurrentPage: page,
			TotalPages:  totalPages,
			TotalItems:  totalCount,
		},
	})
}

// GetUserByID retrieves a user by ID
func (uc *UserController) GetUserByID(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := uc.Service.GetUserByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(utils.Response{
			Status:  "error",
			Message: "User not found",
		})
	}
	return c.Status(fiber.StatusOK).JSON(utils.Response{
		Status: "success",
		Data:   user,
	})
}

// UpdateUser updates an existing user
func (uc *UserController) UpdateUser(c *fiber.Ctx) error {
	idStr := c.Params("id") // Get the ID from the URL parameter
	user := new(models.User)

	// Parse the request body into the user struct
	if err := c.BodyParser(user); err != nil {
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
	user.ID = id // Set the ID for updating

	// Validate the user struct
	if err := uc.Validator.Struct(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.Response{
			Status:  "error",
			Message: utils.ParseValidationErrors(err.(validator.ValidationErrors)),
		})
	}

	// Call the service to update the user
	if err := uc.Service.UpdateUser(user); err != nil {
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

// DeleteUser deletes a user by ID
func (uc *UserController) DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	// log the id
	log.Println(id)
	if err := uc.Service.DeleteUser(id); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(utils.Response{
			Status:  "error",
			Message: "User not found or delete failed",
		})
	}
	return c.Status(fiber.StatusOK).JSON(utils.Response{
		Status:  "success",
		Message: "User deleted successfully",
	})
}

// RegisterUser handles user registration
func (uc *UserController) RegisterUser(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.Response{
			Status:  "error",
			Message: "Invalid input",
		})
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.Response{
			Status:  "error",
			Message: "Failed to hash password",
		})
	}
	user.Password = string(hashedPassword)

	// if role is not provided, set it to "user"
	if user.RoleID == 0 {
		user.RoleID = 2
	}

	if err := uc.Service.CreateUser(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.Response{
			Status:  "error",
			Message: "Failed to create user",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(utils.Response{
		Status:  "success",
		Message: "User registered successfully",
		Data:    user,
	})
}

// LoginUser handles user login
func (uc *UserController) LoginUser(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.Response{
			Status:  "error",
			Message: "Invalid input",
		})
	}

	// Fetch user from the database
	storedUser, err := uc.Service.GetUserByEmail(user.Email)
	if err != nil || storedUser == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.Response{
			Status:  "error",
			Message: "Invalid email or password",
		})
	}

	// Check password
	if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.Response{
			Status:  "error",
			Message: "Invalid email or password",
		})
	}

	// Generate token
	token, err := utils.GenerateToken(storedUser.Email, storedUser.Role.Name)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.Response{
			Status:  "error",
			Message: "Could not generate token",
		})
	}

	return c.Status(fiber.StatusOK).JSON(utils.Response{
		Status: "success",
		Data:   fiber.Map{"token": token, "role": storedUser.Role.Name},
	})
}
