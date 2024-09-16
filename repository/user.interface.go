package repository

import "github.com/alibek-dzhukaev/go-abs-beg/model"

type UserInterface interface {
	CreateUser(model.User) (string, error)
	GetUserById(string) (model.User, error)
	GetAllUsers() ([]model.User, error)
	UpdateUserAgeById(string, int) (int, error)
	DeleteUserById(string) (int, error)
	DeleteAllUsers() (int, error)
}
