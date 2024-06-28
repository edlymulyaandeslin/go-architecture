package repository

import (
	"clean-code-app-laundry/model"
	"database/sql"
	"time"
)

type BillRepository interface {
	Create(payload model.Bill) (model.Bill, error)
}

type billRepository struct {
	db *sql.DB
}

func (b *billRepository) Create(payload model.Bill) (model.Bill, error) {
	transaction, err := b.db.Begin()
	if err != nil {
		return model.Bill{}, err
	}

	var bil model.Bill
	err = transaction.QueryRow("INSERT INTO bills (bill_date, customer_id, user_id, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id, bill_date", time.Now(), payload.Customer.Id, payload.User.Id, time.Now(), time.Now()).Scan(&bil.Id, &bil.BillDate)

	if err != nil {
		return model.Bill{}, transaction.Rollback()
	}

	var billDetails []model.BillDetail
	for _, bd := range payload.BillDetails {
		var billDetail model.BillDetail
		err = transaction.QueryRow("INSERT INTO bill_details (bill_id, product_id, qty, price, created_at, updated_at) VALUES($1, $2, $3, $4, $5, $6) RETURNING id, bill_id, qty, price", bil.Id, bd.Product.Id, bd.Qty, bd.Price, time.Now(), time.Now()).Scan(&billDetail.Id, &billDetail.BillId, &billDetail.Qty, &billDetail.Price)

		if err != nil {
			return model.Bill{}, transaction.Rollback()
		}

		billDetail.Product = bd.Product
		billDetails = append(billDetails, billDetail)
	}

	bil.Customer = payload.Customer
	bil.User = payload.User
	bil.BillDetails = billDetails
	if err = transaction.Commit(); err != nil {
		return model.Bill{}, err
	}

	return bil, err
}

func NewBillRepository(database *sql.DB) BillRepository {
	return &billRepository{db: database}
}
