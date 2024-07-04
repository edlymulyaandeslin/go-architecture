package service

import (
	"clean-code-app-laundry/model"
	"clean-code-app-laundry/model/dto"
	"clean-code-app-laundry/repository"
)

type BillService interface {
	CreateNewBill(payload dto.BillRequest) (model.Bill, error)
}

type billService struct {
	repo            repository.BillRepository
	userService     UserService
	productService  ProductService
	customerService CustomerService
}

func (b *billService) CreateNewBill(payload dto.BillRequest) (model.Bill, error) {
	customer, err := b.customerService.FindById(payload.CustomerId)
	if err != nil {
		return model.Bill{}, err
	}

	user, err := b.userService.FindById(payload.UserId)
	if err != nil {
		return model.Bill{}, err
	}

	var billDetails []model.BillDetail
	for _, bd := range payload.BillDetails {
		product, err := b.productService.FindById(bd.Product.Id)
		if err != nil {
			return model.Bill{}, err
		}

		billDetails = append(billDetails, model.BillDetail{
			Product: product,
			Qty:     bd.Qty,
			Price:   product.Price})
	}

	newPayload := model.Bill{
		Customer:    customer,
		User:        user,
		BillDetails: billDetails,
	}

	bill, err := b.repo.Create(newPayload)
	if err != nil {
		return model.Bill{}, err
	}

	return bill, err
}

func NewBillService(repo repository.BillRepository, us UserService, ps ProductService, cs CustomerService) BillService {
	return &billService{
		repo:            repo,
		userService:     us,
		productService:  ps,
		customerService: cs}
}
