package usermodel

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

// Cascade -> jika parent di hapus maka child akan ikut terhapus
// SET NULL -> jika parent di hapus maka child tetap ada
