package models

import (
	"github.com/jinzhu/gorm"
)

type Users struct {
	gorm.Model
	Nama     string `json:"nama" form:"nama" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required"`
	Role     string `json:"role" form:"role" gorm:"type:enum('Admin', 'User');default:'User'; not-null"`
	Status   string `json:"status" form:"status" gorm:"type:enum('0', '1');default:'0'; not-null"`
}

type Users Login struct {
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required"`
}

// For Response
type UsersResponse struct {
	gorm.Model
	Nama  string `json:"nama" form:"nama" validate:"required"`
	Email string `json:"email" form:"email"`
}

// For JWT Only
type UsersResponses struct {
	ID    uint   `json:"id" form:"id"`
	Nama  string `json:"nama" form:"nama"`
	Email string `json:"email" form:"email"`
	Token string `json:"token" form:"token"`
}
