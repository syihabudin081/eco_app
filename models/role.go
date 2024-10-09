package models // models/role.go

// Role represents a user role in the system
type Role struct {
	ID   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"not null;unique"` // Store role name, e.g., "admin", "user", "guest"
}

// Define available roles as constants for convenience
const (
	AdminRoleName = "admin"
	UserRoleName  = "user"
	GuestRoleName = "guest"
)
