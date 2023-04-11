package services

import (
	"Go_Auliani-Sukma_Auliani-Sukma-Auliani-Sukma/praktikum/models"
	"Go_Auliani-Sukma_Auliani-Sukma-Auliani-Sukma/praktikum/repositories"
)

type UserService interface {
	GetUsersService() ([]*model.User, error)
	CreateService(user model.User) (*model.User, error)
}

type userService struct {
	userR repositories.UserRepository
}

func NewUserService(userR repositories.UserRepository) UserService {
	return &userService{
		userR: userR,
	}
}

func (u *userService) GetUsersService() ([]*model.User, error) {
	users, err := u.userR.GetUsersRepository()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (u *userService) CreateService(user model.User) (*model.User, error) {
	userR, err := u.userR.CreateRepository(user)
	if err != nil {
		return nil, err
	}

	return userR, nil
}