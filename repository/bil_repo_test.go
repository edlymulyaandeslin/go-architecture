package repository

// go test ./... -v -coverprofile cover.out; go tool cover -html cover.out
// magic word

import (
	"clean-code-app-laundry/model"
	"database/sql"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type BillRepositoryTestSuite struct {
	suite.Suite
	mockDb  *sql.DB
	mockSql sqlmock.Sqlmock
	repo    BillRepository
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
				Id: "1",
			},
			Qty:   1,
			Price: 23444,
		},
	},
	CreatedAt: time.Now(),
	UpdatedAt: time.Now(),
}

func (suite *BillRepositoryTestSuite) SetupTest() {
	db, mock, err := sqlmock.New()
	assert.NoError(suite.T(), err)
	suite.mockDb = db
	suite.mockSql = mock
	suite.repo = NewBillRepository(suite.mockDb)
}

func TestBillRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(BillRepositoryTestSuite))
}

func (suite *BillRepositoryTestSuite) TestCreate_Success() {
	suite.mockSql.ExpectBegin()
	rows := sqlmock.NewRows([]string{"id", "bill_date"}).AddRow(mockingBill.Id, mockingBill.BillDate)
	suite.mockSql.ExpectQuery("INSERT INTO bills").WillReturnRows(rows)

	for _, mb := range mockingBill.BillDetails {
		rows := sqlmock.NewRows([]string{"id", "bill_id", "qty", "price"}).AddRow(mb.Id, mb.BillId, mb.Qty, mb.Price)
		suite.mockSql.ExpectQuery("INSERT INTO bill_details").WillReturnRows(rows)
	}

	suite.mockSql.ExpectCommit()
	actual, err := suite.repo.Create(mockingBill)
	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), mockingBill.Id, actual.Id)
}

func (suite *BillRepositoryTestSuite) TestCreate_Failed() {
	suite.mockSql.ExpectBegin().WillReturnError(errors.New("error Begin"))

	_, err := suite.repo.Create(mockingBill)
	assert.Error(suite.T(), err)
}

func (suite *BillRepositoryTestSuite) TestCreateInsertBill_Failed() {
	suite.mockSql.ExpectBegin()

	suite.mockSql.ExpectQuery("INSERT INTO bills").WillReturnError(errors.New("Insert Bill Failed"))
	_, err := suite.repo.Create(mockingBill)
	assert.Error(suite.T(), err)
}

func (suite *BillRepositoryTestSuite) TestCreateInsertBillDetail_Failed() {
	suite.mockSql.ExpectBegin()

	rows := sqlmock.NewRows([]string{"id", "bill_date"}).AddRow(mockingBill.Id, mockingBill.BillDate)
	suite.mockSql.ExpectQuery("INSERT INTO bills").WillReturnRows(rows)

	for _, mb := range mockingBill.BillDetails {
		fmt.Println(mb)
		suite.mockSql.ExpectQuery("INSERT INTO bill_details").WillReturnError(errors.New("Insert Bill Detail Failed"))
		_, err := suite.repo.Create(mockingBill)
		assert.Error(suite.T(), err)
	}
}

func (suite *BillRepositoryTestSuite) TestCreate_FailedCommit() {
	suite.mockSql.ExpectBegin()
	rows := sqlmock.NewRows([]string{"id", "bill_date"}).AddRow(mockingBill.Id, mockingBill.BillDate)
	suite.mockSql.ExpectQuery("INSERT INTO bills").WillReturnRows(rows)

	for _, mb := range mockingBill.BillDetails {
		rows := sqlmock.NewRows([]string{"id", "qty", "price"}).AddRow(mb.Id, mb.Qty, mb.Price)
		suite.mockSql.ExpectQuery("INSERT INTO bill_details").WillReturnRows(rows)
	}

	suite.mockSql.ExpectCommit().WillReturnError(errors.New("commit failed"))
	_, err := suite.repo.Create(mockingBill)
	assert.Error(suite.T(), err)
}
