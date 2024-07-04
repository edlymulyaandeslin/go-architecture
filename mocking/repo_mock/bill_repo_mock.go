package repomock

import (
	"clean-code-app-laundry/model"

	"github.com/stretchr/testify/mock"
)

type BillRepoMock struct {
	mock.Mock
}

func (u *BillRepoMock) Create(payload model.Bill) (model.Bill, error) {
	args := u.Called(payload)
	return args.Get(0).(model.Bill), args.Error(1)
}
