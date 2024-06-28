package service

import (
	"clean-code-app-laundry/model"
	"clean-code-app-laundry/model/dto"
	"clean-code-app-laundry/repository"
	"fmt"
)

// inteface
type ProductService interface {
	FindById(id string) (model.Product, error)
	FindAll(page int, size int) ([]model.Product, dto.Paginate, error)
}

// struct
type productService struct {
	repo repository.ProductRepository
}

func (c *productService) FindById(id string) (model.Product, error) {
	product, err := c.repo.GetById(id)

	if err != nil {
		return model.Product{}, fmt.Errorf("product not found")
	}

	return product, nil
}

func (c *productService) FindAll(page int, size int) ([]model.Product, dto.Paginate, error) {
	return c.repo.GetAll(page, size)
}

// constructor
func NewProductService(repository repository.ProductRepository) ProductService {
	return &productService{repo: repository}
}
