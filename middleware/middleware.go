package middleware

import (
	"belajar-go-fiber/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"strings"
)

// AuthMiddleware ensures the request is authenticated
func AuthMiddleware(c *fiber.Ctx) error {
	// Extract token from header
	tokenString := c.Get("Authorization")
	if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Missing or malformed JWT",
		})
	}
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	// Validate token
	claims, err := utils.ValidateToken(tokenString)
	if err != nil {
		log.Printf("Token validation error: %v\n", err)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Set claims in the context for future use
	c.Locals("userClaims", claims)
	return c.Next()
}

// AdminMiddleware ensures the user has an admin role
func AdminMiddleware(c *fiber.Ctx) error {
	// Check the user role from the JWT claims
	claims, ok := c.Locals("userClaims").(jwt.MapClaims)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "User claims not found",
		})
	}

	role := claims["role"].(string)

	if role != "admin" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "Access restricted to admins only",
		})
	}

	return c.Next()
}
