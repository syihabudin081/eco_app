package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

// GormModel is a model that contains common columns for all tables.
type GormModel struct {
	ID        uuid.UUID      `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
