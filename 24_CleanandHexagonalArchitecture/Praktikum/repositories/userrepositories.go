package repositories

import (
	"Go_Auliani-Sukma_Auliani-Sukma-Auliani-Sukma/praktikum/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUsersRepository() ([]*model.User, error)
	CreateRepository(user model.User) (*model.User, error)
}

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) UserRepository {
	return &userRepository{
		DB: DB,
	}
}

func (u *userRepository) GetUsersRepository() ([]*model.User, error) {
	var users []*model.User

	if err := u.DB.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (u *userRepository) CreateRepository(user model.User) (*model.User, error) {
	if err := u.DB.Save(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}