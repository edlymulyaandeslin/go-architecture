package repository

import (
	"clean-code-app-laundry/model"
	"database/sql"
)

// buat interface
type CustomerRepository interface {
	GetAll(page int, size int) ([]model.Customer, error)
	GetById(id string) (model.Customer, error)
}

// struct
type customerRepository struct {
	db *sql.DB
}

func (p *customerRepository) GetAll(page int, size int) ([]model.Customer, error) {
	panic("unimplemented")
}

func (p *customerRepository) GetById(id string) (model.Customer, error) {
	var customer model.Customer

	err := p.db.QueryRow("SELECT id, name, phone_number, address, created_at, updated_at FROM customers WHERE id=$1", id).Scan(&customer.Id, &customer.Name, &customer.PhoneNumber, &customer.Address, &customer.CreatedAt, &customer.UpdatedAt)

	if err != nil {
		return model.Customer{}, err
	}

	return customer, nil
}

// constructor
func NewCustomerRepository(database *sql.DB) CustomerRepository {
	return &customerRepository{db: database}
}
