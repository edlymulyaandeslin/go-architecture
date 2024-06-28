package controller

import (
	"clean-code-app-laundry/service"
	"clean-code-app-laundry/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	service service.ProductService
	rg      *gin.RouterGroup
}

func (p *ProductController) getAllHandler(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, err2 := strconv.Atoi(c.DefaultQuery("size", "10"))

	if err != nil || err2 != nil {
		util.SendErrorResponse(c, err.Error(), http.StatusBadRequest)
	}

	data, paginate, err := p.service.FindAll(page, size)
	if err != nil {
		util.SendErrorResponse(c, err.Error(), http.StatusInternalServerError)
	}

	var listData []any
	for _, product := range data {
		listData = append(listData, product)
	}

	util.SendPaginateResponse(c, "Success Get Data", listData, paginate, http.StatusOK)
}

func (p *ProductController) Route() {
	router := p.rg.Group("/products")
	router.GET("/", p.getAllHandler)
}

func NewProductController(service service.ProductService, rg *gin.RouterGroup) *ProductController {
	return &ProductController{
		service: service,
		rg:      rg,
	}
}
