package controller

import (
	"clean-code-app-laundry/model/dto"
	"clean-code-app-laundry/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BillController struct {
	service service.BillService
	rg      *gin.RouterGroup
}

func (b *BillController) CreateHandler(c *gin.Context) {
	var payload dto.BillRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	response, err := b.service.CreateNewBill(payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusCreated, response)
}

func (b *BillController) Route() {
	group := b.rg.Group("/transaction")
	group.POST("/", b.CreateHandler)
}

func NewBillController(service service.BillService, rg *gin.RouterGroup) *BillController {
	return &BillController{service: service, rg: rg}
}
