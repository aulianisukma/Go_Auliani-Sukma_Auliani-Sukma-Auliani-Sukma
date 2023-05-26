package models

import "github.com/jinzhu/gorm"

type Kategori struct {
	gorm.Model
	Nama	string 	`json:"nama" form:"nama"`
}