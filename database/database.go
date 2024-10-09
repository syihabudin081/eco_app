package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB is the database connection
var DB *gorm.DB

// Connect to the database
func DatabaseInit() {
	var err error
	// Ganti dengan konfigurasi database Anda
	dsn := "host=localhost user=postgres password=123 dbname=eco_db port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	fmt.Println("Database connection established")
}
