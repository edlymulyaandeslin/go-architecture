package controller

import (
	"bytes"
	"clean-code-app-laundry/middleware"
	"clean-code-app-laundry/mocking"
	servicemock "clean-code-app-laundry/mocking/service_mock"
	"clean-code-app-laundry/model"
	"clean-code-app-laundry/model/dto"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type BillControllerTestSuite struct {
	suite.Suite
	serviceBillMock *servicemock.BillServiceMock
	rg              *gin.RouterGroup
	middlewareMock  middleware.AuthMiddleware
}

func (suite *BillControllerTestSuite) SetupTest() {
	suite.serviceBillMock = new(servicemock.BillServiceMock)
	rg := gin.Default()
	suite.rg = rg.Group("/api/v1/transaction")
	suite.middlewareMock = new(mocking.AuthMiddlewareMock)
}

func TestBillController(t *testing.T) {
	suite.Run(t, new(BillControllerTestSuite))
}

var mockingBill = model.Bill{
	Id:       "1",
	BillDate: time.Now(),
	Customer: model.Customer{
		Id:   "1",
		Name: "Arfian",
	},
	User: model.User{
		Id:   "2",
		Name: "Dimas",
	},
	BillDetails: []model.BillDetail{
		{
			Id:     "1",
			BillId: "1",
			Product: model.Product{
				Id:    "1",
				Price: 1,
			},
			Qty:   1,
			Price: 1,
		},
	},
	CreatedAt: time.Now(),
	UpdatedAt: time.Now(),
}

var mockPayload = dto.BillRequest{
	CustomerId: "1",
	UserId:     "1",
	BillDetails: []model.BillDetail{
		{
			Product: model.Product{
				Id: "1",
			},
			Qty: 1,
		},
	},
}

func (suite *BillControllerTestSuite) TestCreateHandler_success() {
	suite.serviceBillMock.On("CreateNewBill", mockPayload).Return(mockingBill, nil)

	record := httptest.NewRecorder()
	mockPayloadJson, err := json.Marshal(mockPayload)
	assert.NoError(suite.T(), err)

	req, err := http.NewRequest(http.MethodPost, "/", bytes.NewBuffer(mockPayloadJson))
	assert.NoError(suite.T(), err)

	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJFbmlnbWFDYW1wIiwiZXhwIjoxNzE5ODMzNTYyLCJpYXQiOjE3MTk4Mjk5NjIsInVzZXJJZCI6IjkxN2RhOWI3LTdhZmMtNDk2Yi04MzNlLTNhN2NlZWNjYmIxYyJ9.40sqR3uRdlfge8MLyhD6XrXHTqMBzLYLc_2o7StmqD4"

	req.Header.Set("Authorization", "Bearer "+token)
	ctx, _ := gin.CreateTestContext(record)
	ctx.Request = req

	suite.serviceBillMock.On("CreateNewBill", mockPayload).Return(mockingBill, nil)

	bilController := NewBillController(suite.serviceBillMock, suite.rg, suite.middlewareMock)

	bilController.Route()
	bilController.CreateHandler(ctx)

	mockBillJson, _ := json.Marshal(mockingBill)

	assert.Equal(suite.T(), http.StatusCreated, record.Code)
	assert.Equal(suite.T(), string(mockBillJson), record.Body.String())
}

func (suite *BillControllerTestSuite) TestCreateHandler_failed() {
	bilController := NewBillController(suite.serviceBillMock, suite.rg, suite.middlewareMock)
	bilController.Route()
	record := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/", nil)
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJFbmlnbWFDYW1wIiwiZXhwIjoxNzE5ODMzNTYyLCJpYXQiOjE3MTk4Mjk5NjIsInVzZXJJZCI6IjkxN2RhOWI3LTdhZmMtNDk2Yi04MzNlLTNhN2NlZWNjYmIxYyJ9.40sqR3uRdlfge8MLyhD6XrXHTqMBzLYLc_2o7StmqD4"

	req.Header.Set("Authorization", "Bearer "+token)
	ctx, _ := gin.CreateTestContext(record)
	ctx.Request = req

	bilController.CreateHandler(ctx)

	assert.Equal(suite.T(), http.StatusBadRequest, record.Code)
}
