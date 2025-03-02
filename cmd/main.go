package main

import (
	"decard/internal/application/query"
	ormRepository "decard/internal/infrastructure/orm/repositoty"
	"decard/internal/infrastructure/provider"
	"decard/internal/infrastructure/provider/api"
	"decard/internal/presentation/http/handlers"
	"log"
	"net/http"

	"decard/config"
	"decard/internal/domain/service"
	"decard/internal/infrastructure/database"
	"decard/internal/presentation/http/routes"
)

func main() {
	// Load configuration
	cfg := config.NewConfig()

	// Initialize database
	db := database.NewDatabase(cfg)
	defer db.Close()

	// Initialize repositories
	ProfileRepository := ormRepository.NewProfileRepository(db.DB)

	providerClient := provider.NewClient(*cfg)

	accountAPI := api.NewAccountApi(providerClient)
	paymentAPI := api.NewPaymentApi(providerClient)
	transactionAPI := api.NewTransactionApi(providerClient)

	// Initialize services
	authService := service.NewAuthService(ProfileRepository)
	accountService := service.NewAccountService(accountAPI)
	paymentService := service.NewPaymentService(paymentAPI)
	transactionService := service.NewTransactionService(transactionAPI)

	//CQRS
	getAccountCardsQuery := query.NewGetAccountCardsHandler(accountService)
	getCardTransactionsQueryHandler := query.NewGetCardTransactionsQueryHandler(transactionService)

	// Initialize handlers
	authHandler := handlers.NewAuthHandler(*authService)
	accountHandler := handlers.NewAccountHandler(*accountService, *getAccountCardsQuery)
	paymentHandler := handlers.NewPaymentHandler(*paymentService)
	transactionHandler := handlers.NewTransactionHandler(getCardTransactionsQueryHandler)

	// Setup router
	router := routes.NewRouter(
		authHandler,
		accountHandler,
		paymentHandler,
		transactionHandler,
	)

	// Start server
	log.Printf("Server starting on %s", cfg.ServerAddress)
	if err := http.ListenAndServe(cfg.ServerAddress, router); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
