package models

import (
	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	CategoryID  int      `json:"category_id" form:"category_id"`
	Category    Category `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Name        string   `json:"name" form:"name"`
	Description string   `json:"description" form:"description"`
	Stock       int      `json:"stock" form:"stock"`
	Price       int      `json:"price" form:"price"`
}
