package servicemock

import (
	"clean-code-app-laundry/model"
	"clean-code-app-laundry/model/dto"

	"github.com/stretchr/testify/mock"
)

type ProductServiceMock struct {
	mock.Mock
}

func (u *ProductServiceMock) FindById(id string) (model.Product, error) {
	args := u.Called(id)
	return args.Get(0).(model.Product), args.Error(1)
}

func (u *ProductServiceMock) FindAll(page int, size int) ([]model.Product, dto.Paginate, error) {
	args := u.Called(page, size)
	return args.Get(0).([]model.Product), args.Get(1).(dto.Paginate), args.Error(2)
}
