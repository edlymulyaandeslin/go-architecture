package service

import (
	"clean-code-app-laundry/model"
	"clean-code-app-laundry/repository"
	"fmt"
)

// inteface
type UserService interface {
	FindById(id string) (model.User, error)
	FindAll(page int, size int) ([]model.User, error)
}

// struct
type userService struct {
	repo repository.UserRepository
}

func (c *userService) FindById(id string) (model.User, error) {
	user, err := c.repo.GetById(id)

	if err != nil {
		return model.User{}, fmt.Errorf("user not found")
	}

	return user, nil
}
func (c *userService) FindAll(page int, size int) ([]model.User, error) {
	panic("unimplement")
}

// constructor
func NewUserService(repository repository.UserRepository) UserService {
	return &userService{repo: repository}
}
