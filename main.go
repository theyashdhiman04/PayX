package main

import (
	"fmt"
	"log"

	"github.com/theyashdhiman04/PayX/api/routes"
	"github.com/theyashdhiman04/PayX/config"
	repository "github.com/theyashdhiman04/PayX/internal/repositories"
	"github.com/theyashdhiman04/PayX/internal/services"
	"github.com/theyashdhiman04/PayX/storage"
)

func main() {
	appConfig, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load configuration: %v", err)
	}

	database, err := storage.NewPostgresDB(appConfig)
	if err != nil {
		log.Fatalf("failed to establish database connection: %v", err)
	}
	fmt.Println("Database connection established successfully")

	userRepo := repository.NewUserRepository(database)
	transactionRepo := repository.NewTransactionRepository(database)

	authSvc := services.NewAuthService(userRepo, appConfig)
	userSvc := services.NewUserService(userRepo)
	midtransSvc := services.NewMidtransService(appConfig)
	paymentSvc := services.NewPaymentService(transactionRepo, midtransSvc)

	router := routes.SetupRouter(authSvc, userSvc, paymentSvc, appConfig)

	serverAddr := fmt.Sprintf(":%s", appConfig.ServerPort)
	log.Printf("Starting server on port %s", appConfig.ServerPort)
	if err := router.Run(serverAddr); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
