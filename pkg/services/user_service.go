package services

import (
	"fiber-apis/pkg/models"
	"fiber-apis/pkg/repositories"
)

type UserServiceInterface interface {
	GetAllUsersService() ([]models.User, error)
	GetUserById(_id int) (models.User, error)
	CreateUser(_newUser models.User) error
	DeleteUserById(_id int) error
	UpdateUserById(_id int, _user models.User) error
}

// var _ UserService = user{}

type UserServiceStruct struct {
	Repository repositories.UserRepository
}

func (s UserServiceStruct) GetAllUsersService() ([]models.User, error) {

	return s.Repository.GetAllUsersRepository()
}

func (s UserServiceStruct) GetUserById(_id int) (models.User, error) {

	return s.Repository.GetUserById(_id)

}

func (s UserServiceStruct) CreateUser(_newUser models.User) error {

	return s.Repository.CreateUser(_newUser)
}

func (s UserServiceStruct) DeleteUserById(_id int) error {

	return s.Repository.DeleteUserById(_id)

}

func (s UserServiceStruct) UpdateUserById(_id int, _user models.User) error {

	return s.Repository.UpdateUserById(_id, _user)
}
