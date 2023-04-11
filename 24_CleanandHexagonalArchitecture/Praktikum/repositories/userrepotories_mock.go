package repositories

import (
	"Go_Auliani-Sukma_Auliani-Sukma-Auliani-Sukma/praktikum/model"

	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock interface {
	GetUsersRepository() ([]*model.User, error)
	CreateRepository(userData model.User) (*model.User, error)
}

type IuserRepositoryMock struct {
	Mock mock.Mock
}

func NewUserRepositoryMock(mock mock.Mock) UserRepositoryMock {
	return &IuserRepositoryMock{
		Mock: mock,
	}
}

func (u *IuserRepositoryMock) GetUsersRepository() ([]*model.User, error) {
	args := u.Mock.Called()
	if args.Get(0) == nil {
		return nil, args.Get(1).(error)
	}

	users := args.Get(0).([]*model.User)

	return users, nil
}

func (u *IuserRepositoryMock) CreateRepository(userData model.User) (*model.User, error) {
	args := u.Mock.Called(userData)
	if args.Get(0) == nil {
		return nil, args.Get(1).(error)
	}

	user := args.Get(0).(model.User)

	return &user, nil
}