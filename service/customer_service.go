package service

import (
	"clean-code-app-laundry/model"
	"clean-code-app-laundry/repository"
	"fmt"
)

// inteface
type CustomerService interface {
	FindById(id string) (model.Customer, error)
	FindAll(page int, size int) ([]model.Customer, error)
}

// struct
type customerService struct {
	repo repository.CustomerRepository
}

func (c *customerService) FindById(id string) (model.Customer, error) {
	customer, err := c.repo.GetById(id)

	if err != nil {
		return model.Customer{}, fmt.Errorf("customer with id %s not found", id)
	}

	return customer, nil
}
func (c *customerService) FindAll(page int, size int) ([]model.Customer, error) {
	panic("unimplement")
}

// constructor
func NewCustomerService(repository repository.CustomerRepository) CustomerService {
	return &customerService{repo: repository}
}
