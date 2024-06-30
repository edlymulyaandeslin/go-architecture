package controller

import (
	"clean-code-app-laundry/model"
	"clean-code-app-laundry/model/dto"
	"clean-code-app-laundry/service"
	"clean-code-app-laundry/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service service.UserService
	rg      *gin.RouterGroup
}

func (u *UserController) registerHandler(c *gin.Context) {
	payload := model.User{}
	if err := c.ShouldBindJSON(&payload); err != nil {
		util.SendErrorResponse(c, "failed to parsing payload", http.StatusInternalServerError)
	}
	data, err := u.service.CreateNew(payload)
	if err != nil {
		util.SendErrorResponse(c, err.Error(), http.StatusInternalServerError)
		return
	}
	util.SendSingleResponse(c, "success create new user", data, http.StatusOK)
}

func (u *UserController) FindByUsernameHandler(c *gin.Context) {
	username := c.Param("username")

	user, err := u.service.FindByUsername(username)
	if err != nil {
		util.SendErrorResponse(c, "error", http.StatusBadRequest)
	}
	util.SendSingleResponse(c, "succes retrieving user", user, http.StatusOK)
}

func (u *UserController) LoginHandler(c *gin.Context) {
	var payload dto.LoginDto
	if err := c.ShouldBindJSON(&payload); err != nil {
		util.SendErrorResponse(c, "Failed to parsing payload", http.StatusBadRequest)
		return
	}

	response, errors := u.service.Login(payload)
	if errors != nil {
		util.SendErrorResponse(c, errors.Error(), http.StatusInternalServerError)
		return
	}
	util.SendSingleResponse(c, "success login", response, http.StatusOK)
}

func (u *UserController) Route() {
	router := u.rg.Group("users")
	router.POST("/register", u.registerHandler)
	router.GET("/:username", u.FindByUsernameHandler)
	router.POST("/login", u.LoginHandler)
}

func NewUserController(service service.UserService, rg *gin.RouterGroup) *UserController {
	return &UserController{service: service, rg: rg}
}
