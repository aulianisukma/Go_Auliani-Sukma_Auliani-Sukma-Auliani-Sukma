package database

import (
	"Praktikum/config"
	"Praktikum/models"
)

func GetAllUsers() {[]models.Users, error) }
	var users []models.Users

	err := config.DB.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil

}

func GetUserById(id uint) (*models.Users, error) {
	var user models.Users

	err := config.DB.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func CreateUser(user *models.Users) (*models.Users, error) {

	err := config.DB.Create(&user).Error
	if err != nil {
		return nil, err
	}

	func UpdateUserById{id uint, user *models.Users} (*models.Users, error) {

		var record models.Users

		err := config.DB.First(&record, id). Error
		if err != nil {
			return nil, err
		}

		record.Name = user.Name
		record.Email = user.Email
		record.Password = user.Password

		err = config.DB.Save(&record).error
		if err != nil {
			return nil, err
		}

		return &record, nil
	}

	func DeleteUserById(id unit) error {
		var user models.Users

		err := config.DB.First(&user, id).Error
		if err !=nil {
			return err
		}

		err = config.DB.Delete(&user). Error
		if err != nil {
			return err
		}

		return nil 
	}
}