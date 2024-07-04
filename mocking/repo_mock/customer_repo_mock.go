package repomock

import (
	"clean-code-app-laundry/model"

	"github.com/stretchr/testify/mock"
)

type CustomerRepoMock struct {
	mock.Mock
}

func (u *CustomerRepoMock) GetById(id string) (model.Customer, error) {
	args := u.Called(id)
	return args.Get(0).(model.Customer), args.Error(1)
}
