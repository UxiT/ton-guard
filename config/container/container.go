package container

import (
	"database/sql"
	"decard/config"
	"decard/internal/application/command/auth"
	"decard/internal/application/query/accounts"
	"decard/internal/application/query/transactions"
	"decard/internal/domain/service"
	"decard/internal/infrastructure/database"
	ormrepository "decard/internal/infrastructure/orm/repositoty"
	"decard/internal/infrastructure/provider"
	"decard/internal/infrastructure/provider/api"
	"decard/internal/presentation/http/handlers"
	"decard/internal/presentation/http/routes"
	"github.com/gorilla/mux"
	"log/slog"
	"os"
)

type Container struct {
	Logger *slog.Logger
	Router *mux.Router
	DB     *sql.DB
}

func NewContainer(cfg *config.Config) *Container {
	logger := setupLogger(cfg.Env)

	//db
	db := database.NewDatabase(cfg.DbUrl)

	//repositories
	profileRepo := ormrepository.NewProfileRepository(db.DB)
	accountRepo := ormrepository.NewAccountRepository(db.DB)
	customerRepo := ormrepository.NewCustomerRepository(db.DB)
	refreshTokenRepo := ormrepository.NewRefreshTokenRepository(db.DB, cfg.RefreshTokenTTL)

	//external client
	providerClient := provider.NewClient(*cfg)

	accountAPI := api.NewAccountApi(providerClient)
	paymentAPI := api.NewPaymentApi(providerClient)
	transactionAPI := api.NewTransactionApi(providerClient)

	//services
	paymentService := service.NewPaymentService(paymentAPI)
	transactionService := service.NewTransactionService(transactionAPI)
	authService := service.NewAuthService(logger, refreshTokenRepo)

	//CQRS
	generateJWTCommand := auth.NewGenerateJWTCommandHandler(logger, authService)
	authCommand := auth.NewAuthenticateCommandHandler(logger, profileRepo, generateJWTCommand)
	registerCommand := auth.NewRegisterCommandHandler(logger, profileRepo, generateJWTCommand)
	refreshAuthCommand := auth.NewRefreshCommandHandler(logger, refreshTokenRepo, generateJWTCommand)

	getAccountCardsQuery := accounts.NewGetAccountCardsHandler(accountAPI)
	getAccountForProfileQuery := accounts.NewGetAccountForProfileQueryHandler(logger, accountAPI, customerRepo, accountRepo)
	getAccountListQuery := accounts.NewGetAccountListQueryHandler(logger, accountAPI)

	getCardTransactionsQueryHandler := transactions.NewGetCardTransactionsQueryHandler(transactionService)

	//handlers
	authHandler := handlers.NewAuthHandler(authCommand, registerCommand, refreshAuthCommand)
	accountHandler := handlers.NewAccountHandler(
		logger,
		getAccountCardsQuery,
		getAccountForProfileQuery,
		getAccountListQuery,
	)
	paymentHandler := handlers.NewPaymentHandler(*paymentService)
	transactionHandler := handlers.NewTransactionHandler(getCardTransactionsQueryHandler)
	cardHandler := handlers.NewCardHandler(logger)

	router := routes.NewRouter(
		logger,
		authHandler,
		accountHandler,
		paymentHandler,
		transactionHandler,
		cardHandler,
	)

	return &Container{
		Logger: logger,
		Router: router,
		DB:     db.DB,
	}
}

const (
	envLocal      = "local"
	envDevelop    = "develop"
	envProduction = "production"
)

func setupLogger(env string) *slog.Logger {
	var logger *slog.Logger

	switch env {
	case envLocal:
		logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envDevelop:
		logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProduction:
		logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	default:
		panic("unknown env: " + env)
	}

	return logger
}
