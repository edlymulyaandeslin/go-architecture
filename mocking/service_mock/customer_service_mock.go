package servicemock

import (
	"clean-code-app-laundry/model"

	"github.com/stretchr/testify/mock"
)

type CustomerServiceMock struct {
	mock.Mock
}

func (u *CustomerServiceMock) FindById(id string) (model.Customer, error) {
	args := u.Called(id)
	return args.Get(0).(model.Customer), args.Error(1)
}
