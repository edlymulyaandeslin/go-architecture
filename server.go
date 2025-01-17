package main

import (
	"clean-code-app-laundry/config"
	"clean-code-app-laundry/controller"
	"clean-code-app-laundry/middleware"
	"clean-code-app-laundry/repository"
	"clean-code-app-laundry/service"
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Server struct {
	bS      service.BillService
	cS      service.CustomerService
	pS      service.ProductService
	uS      service.UserService
	jS      service.JwtService
	aM      middleware.AuthMiddleware
	engine  *gin.Engine
	portApp string
}

func (s *Server) initiateRoute() {
	routeGroup := s.engine.Group("/api/v1")
	fmt.Println("eror disini")
	controller.NewBillController(s.bS, routeGroup, s.aM).Route()
	controller.NewProductController(s.pS, routeGroup, s.aM).Route()
	controller.NewUserController(s.uS, routeGroup).Route()
}

func (s *Server) Start() {
	fmt.Println("start eror disini")
	s.initiateRoute()
	s.engine.Run(s.portApp)
}

func NewServer() *Server {

	cf, err := config.NewConfig()

	urlConnection := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", cf.Host, cf.Port, cf.User, cf.Password, cf.DbName)

	db, err := sql.Open(cf.Driver, urlConnection)
	if err != nil {
		log.Fatal(err)
	}

	portApp := cf.AppPort
	if portApp == "" {
		portApp = "8080"
	}

	billRepo := repository.NewBillRepository(db)
	custRepo := repository.NewCustomerRepository(db)
	productRepo := repository.NewProductRepository(db)
	userRepo := repository.NewUserRepository(db)

	jwtService := service.NewJwtService(cf.SecurityConfig)
	cusService := service.NewCustomerService(custRepo)
	userService := service.NewUserService(userRepo, jwtService)
	productService := service.NewProductService(productRepo)
	billService := service.NewBillService(billRepo, userService, productService, cusService)

	authMiddleware := middleware.NewAuthMiddleware(jwtService)

	return &Server{
		bS:      billService,
		cS:      cusService,
		pS:      productService,
		uS:      userService,
		jS:      jwtService,
		aM:      authMiddleware,
		engine:  gin.Default(),
		portApp: portApp,
	}

}
