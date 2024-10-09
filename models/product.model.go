package models

// Product represents a product in the system
type Product struct {
	GormModel
	Name        string `json:"name" gorm:"not null"`
	Brand       string `json:"brand" gorm:"not null"`
	Description string `json:"description" gorm:"not null"`
	Eco_Score   string `json:"eco_score" gorm:"not null"`
	Category    string `json:"category" gorm:"not null"`
	Certificate string `json:"certificate" gorm:"not null"`
	Price       int    `json:"price" gorm:"not null"`
}
