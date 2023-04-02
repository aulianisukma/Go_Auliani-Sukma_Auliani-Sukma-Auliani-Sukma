package models 

import (
	"time"

	"gorm.io/gorm"
)

type User Struct {
	ID uint 'gorm:"primaryKey"'
	Name string 'json:"name" form:"name"'
	Email string 'json:"email" form:"email"'
	Password string 'json:"Password" form: UserID'
	DeletedAt gorm.DeletedAt
	CreatedAt time.Time
	UpdatedAt time.Time
}