package main

import (
	"log"

	"github.com/ahmad20/bri-mini-project/configs"
	"github.com/ahmad20/bri-mini-project/modules/account"
	"github.com/ahmad20/bri-mini-project/modules/auth"
	"github.com/ahmad20/bri-mini-project/modules/customer"
	"github.com/ahmad20/bri-mini-project/pkg/database"
	"github.com/ahmad20/bri-mini-project/repositories"
	"github.com/gin-gonic/gin"
)

func main() {
	config, err := configs.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load configuration: %v", err)
	}

	err = database.ConnectDB(config.Database.Host, config.Database.User, config.Database.Password, config.Database.Name, config.Database.Port)
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}
	router := gin.Default()

	authService := auth.NewService(config.JWT.SecretKey, config.JWT.ExpiresIn)
	LoginMiddleware := auth.AuthMiddleware(authService)
	accountRepo := repositories.NewAccountRepository(database.DB)
	accountUseCase := account.NewUseCase(accountRepo)
	customerRepo := repositories.NewCustomerRepository(database.DB)
	customerUseCase := customer.NewUseCase(customerRepo)
	accountHandler := account.NewHandler(accountUseCase, customerUseCase, authService)
	StatusMinddleware := auth.StatusAuthorization(accountRepo)
	account.SetupRouter(router, accountHandler, LoginMiddleware, StatusMinddleware)
	router.Run()
}
