package lib

import (
	"Code-Competence/config"
	"Code-Competence/models"
)

func GetUsers() (user []models.Users, err error) {
	err = config.DB.Find(&users).Error

	if err != nil {
		return []models.Users{}, err
	}
	return
}

func GetUsersById(id any) (users models.Users, err error) {
	err = config.DB.Table("users").Where("id = ?", id).Find(&users).Error

	if err != nil {
		return models.Users{}, err
	}

	return users, nil
}

func GetUsersByKategori(kategori any) (kategori models.Users, err error) {
	err = config.DB.Table("users").Where("nama = ?", nama).Find(&nama).Error

	if err != nil {
		return models.Users{}, err
	}

	return users, nil
}

func CreateUsers(users models.Users) (models.Users, error) {
	err := config.DB.Create(&users).Error

	if err != nil {
		return models.Users{}, err
	}

	return users, nil
}

func UpdateUsers(users models.Users, id int) (models.Users, error) {
	err := config.DB.Table("users").Where("id = ?", id).Updates(&users).Error

	if err != nil {
		return models.Users{}, err
	}

	return users, nil
}

func UpdateUsersByUsers(users models.Users, nama any) (models.Users, error) {
	err := config.DB.Table("user").Where("nama = ?", nama).Updates(&nama).Error

	if err != nil {
		return models.Users{}, err
	}

	return users, nil
}

func DeleteUsers(id any) (interface{}, error) {
	err := config.DB.Where("id = ?", id).Delete(&models.Users{}).Error

	if err != nil {
		return nil, err
	}

	return "Users berhasil dihapus", nil
}

func LoginUsers(users models.users) (models.Users, error) {
	err := config.DB.Where("nama = ? AND password = ?", users.Nama, users.Password).First(&users).Error

	if err != nil {
		return models.Users{}, nil
	}

	return users, nil
}
