package service

import (
	"clean-code-app-laundry/model"
	"clean-code-app-laundry/model/dto"
	"clean-code-app-laundry/repository"
	"clean-code-app-laundry/util"
	"errors"
	"fmt"
)

// inteface
type UserService interface {
	FindById(id string) (model.User, error)
	CreateNew(payload model.User) (model.User, error)
	FindByUsername(username string) (model.User, error)
	Login(payload dto.LoginDto) (dto.LoginResponseDto, error)
}

// struct
type userService struct {
	repo repository.UserRepository
	jwt  JwtService
}

func (c *userService) FindById(id string) (model.User, error) {
	user, err := c.repo.GetById(id)

	if err != nil {
		return model.User{}, fmt.Errorf("user not found")
	}

	return user, nil
}

func (c *userService) CreateNew(payload model.User) (model.User, error) {
	if !payload.IsValidRole() {
		return model.User{}, errors.New("role is invalid, must be admin or employee")
	}
	passwordHash, err := util.EncryptPassword(payload.Password)

	if err != nil {
		return model.User{}, err
	}
	payload.Password = passwordHash
	return c.repo.CreateUser(payload)
}

func (c *userService) FindByUsername(username string) (model.User, error) {
	user, err := c.repo.FindByUsername(username)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (c *userService) Login(payload dto.LoginDto) (dto.LoginResponseDto, error) {
	user, err := c.repo.FindByUsername(payload.Username)
	if err != nil {
		return dto.LoginResponseDto{}, fmt.Errorf("username or password invalid ")
	}

	err = util.ComparePasswordHash(user.Password, payload.Password)
	if err != nil {
		fmt.Println(user.Password, payload.Password)
		return dto.LoginResponseDto{}, fmt.Errorf("password incorrect!")
	}

	user.Password = ""
	token, err := c.jwt.GenerateToken(user)
	if err != nil {
		return dto.LoginResponseDto{}, fmt.Errorf("failed to create token!")
	}

	return token, nil
}

// constructor
func NewUserService(repositori repository.UserRepository, jwt JwtService) UserService {
	return &userService{repo: repositori, jwt: jwt}
}
