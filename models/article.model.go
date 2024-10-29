package models

type Article struct {
	GormModel
	Title    string `json:"title" gorm:"not null"`
	ImageURL string `json:"image_url" gorm:"not null"`
	Content  string `json:"content" gorm:"not null"`
	Author   string `json:"author" gorm:"not null"`
}
