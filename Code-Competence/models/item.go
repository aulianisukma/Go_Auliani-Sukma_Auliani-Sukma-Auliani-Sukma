package models

import "github.com/jinzhu/gorm"

type Item struct {
	gorm.Model
	Nama         string `json:"nama" form:"nama" validate:"required"`
	Deskripsi    string `json:"Deskripsi" form:"nama" validate:"required"`
	Jumlah_Stok  string `json:"jumlah_stok" form:"jumlah_stok" validate:"required"`
	Harga        int    `json:"harga" form:"harga" validate:"required"`
	Kategori     int    `json:"kategori" form:"kategori" validate:"required"`
}

// For Response Item
type ItemResponse struct {
	Nama        string `json:"nama" form:"nama" validate:"required"`
	Deskripsi   string `json:"deskripsi" form:"deskripsi" validate:"required"`
	Jumlah_Stok string `json:"jumlah_stok" form:"jumlah_stok" validate:"required"`
	Harga       int    `json:"harga" form:"harga" validate:"required"`
	Kategori    int    `json:"kategori" form:"kategori" validate:"required"`
}
