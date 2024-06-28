package dto

import "clean-code-app-laundry/model"

type BillRequest struct {
	Id          string             `json:"id"`
	CustomerId  string             `json:"customerId"`
	UserId      string             `json:"userId"`
	BillDetails []model.BillDetail `json:"billDetails"`
}
