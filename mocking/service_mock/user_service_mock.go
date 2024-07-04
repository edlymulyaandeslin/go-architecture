package servicemock

import (
	"clean-code-app-laundry/model"
	"clean-code-app-laundry/model/dto"

	"github.com/stretchr/testify/mock"
)

type UserServiceMock struct {
	mock.Mock
}

func (u *UserServiceMock) FindById(id string) (model.User, error) {
	args := u.Called(id)
	return args.Get(0).(model.User), args.Error(1)
}

func (u *UserServiceMock) CreateNew(payload model.User) (model.User, error) {
	args := u.Called(payload)
	return args.Get(0).(model.User), args.Error(1)
}

func (u *UserServiceMock) FindByUsername(username string) (model.User, error) {
	args := u.Called(username)
	return args.Get(0).(model.User), args.Error(1)
}

func (u *UserServiceMock) Login(payload dto.LoginDto) (dto.LoginResponseDto, error) {
	args := u.Called(payload)
	return args.Get(0).(dto.LoginResponseDto), args.Error(1)
}
