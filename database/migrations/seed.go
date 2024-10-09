package migrations

import (
	"belajar-go-fiber/models"
	"gorm.io/gorm"
)

// SeedRoles will populate the roles table with default values
func SeedRoles(db *gorm.DB) {
	roles := []models.Role{
		{Name: "admin"},
		{Name: "user"},
		{Name: "guest"},
	}

	for _, role := range roles {
		if err := db.FirstOrCreate(&role, models.Role{Name: role.Name}).Error; err != nil {
			panic(err) // Handle the error appropriately in a real app
		}
	}
}
