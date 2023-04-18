package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	ID       int    `json:"id" form:"name"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}
type Token struct {
	Token string `json:"token" form:"token"`
}
