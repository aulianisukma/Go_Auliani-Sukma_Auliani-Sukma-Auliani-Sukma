package database

import (
	"errors"
	middleware "Praktikum/Middlewares"
	"Praktikum/config"
	"Praktikum/models"
)

func GetDetailUsers(userId int) (interface{}, error) {
	var user models.User

	if e := config.DB.Find(&user, userId).Error; e != nil {
		return nil, e
	}
	return user, nil
}

func LoginUsers(user *models.User) (interface{}, error) {
	var err error
	if err = config.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(user).Error; err != nil {
		return nil, err
	}

	var token models.Token
	token.Token, err = middleware.CreateJWTToken(user.ID, user.Name)

	if err != nil {
		return nil, errors.New("failed to create JWT token")
	}

	if err := config.DB.Save(user).Error; err != nil {
		return nil, errors.New("failed to save user to database")
	}

	return token, nil
}
