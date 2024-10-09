package models

type User struct {
	GormModel
	Name     string `json:"name" form:"name" validate:"gte=6,lte=32" gorm:"not null"`
	Email    string `json:"email" form:"email" validate:"required,email" gorm:"not null;unique"`         // Email is unique
	Password string `json:"password" form:"password" validate:"required,gte=8" gorm:"not null"`          // Remove `column:password` as it’s incorrect syntax
	Phone    int    `json:"phone" form:"phone" validate:"required,number,min=12" gorm:"not null;unique"` // Phone is unique
	RoleID   int    `json:"role_id"`                                                                     // Foreign key reference to Role
	Role     Role   `json:"role" gorm:"foreignKey:RoleID"`                                               // Include Role information
}
