package domain

import (
	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Email    string    `gorm:"uniqueIndex;not null"`
	Name     string    `gorm:"type:varchar(255);not null"`
	Phone    string    `gorm:"type:varchar(127);uniqueIndex;not null"`
	Password string    `gorm:"not null"`
	Role     string    `gorm:"type:varchar(255);not null"`
}

type SignUpInput struct {
	Email           string `json:"email" binding:"required"`
	Name            string `json:"name" binding:"required"`
	Phone           string `json:"phone" binding:"required"`
	Password        string `json:"password" binding:"required,min=8"`
	PasswordConfirm string `json:"passwordConfirm" binding:"required,min=8"`
	Role            string `json:"role" binding:"required"`
}

type SignInInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required,min=8"`
}

type SignInResponse struct {
	Token string `json:"token"`
}
