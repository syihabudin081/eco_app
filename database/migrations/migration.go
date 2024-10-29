package migrations

import (
	"belajar-go-fiber/database"
	"belajar-go-fiber/models"
	"fmt"
	"log"
)

func Migration() {
	err := database.DB.AutoMigrate(
		&models.User{},
		&models.Product{},
		&models.Article{},
	)
	if err != nil {
		log.Fatal("Failed to migrate...")
	}

	// Seed roles
	SeedRoles(database.DB)

	fmt.Println("Migrated successfully")
}
