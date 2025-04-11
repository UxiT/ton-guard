package container

import (
	"database/sql"
	"decard/config"
	"decard/internal/application/command/auth"
	"decard/internal/application/query/accounts"
	"decard/internal/application/query/card"
	"decard/internal/application/query/transactions"
	"decard/internal/domain/service"
	"decard/internal/infrastructure/database"
	ormrepository "decard/internal/infrastructure/orm/repositoty"
	"decard/internal/infrastructure/provider"
	"decard/internal/infrastructure/provider/api"
	"decard/internal/presentation/http/handlers/acoount"
	auth2 "decard/internal/presentation/http/handlers/auth"
	card2 "decard/internal/presentation/http/handlers/card"
	"decard/internal/presentation/http/handlers/operation"
	"decard/internal/presentation/http/handlers/payment"
	"decard/internal/presentation/http/handlers/transaction"
	"decard/internal/presentation/http/routes"
	"decard/pkg/utils/decryptor"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	"os"
)

type Container struct {
	Logger *zerolog.Logger
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
	providerClient := provider.NewClient(*cfg, logger)

	accountAPI := api.NewAccountApi(logger, providerClient)
	paymentAPI := api.NewPaymentApi(providerClient)
	transactionAPI := api.NewTransactionApi(providerClient)
	cardAPI := api.NewCardAPI(providerClient, logger, cfg.PublicKey, cfg.PrivateKey)

	//services
	paymentService := service.NewPaymentService(paymentAPI)
	transactionService := service.NewTransactionService(transactionAPI)
	authService := service.NewAuthService(logger, refreshTokenRepo)

	decryptService := decryptor.NewDecryptor(cfg.PrivateKey, logger)

	//CQRS
	generateJWTCommand := auth.NewGenerateJWTCommandHandler(logger, authService)
	authCommand := auth.NewAuthenticateCommandHandler(logger, profileRepo, generateJWTCommand)
	registerCommand := auth.NewRegisterCommandHandler(logger, profileRepo, generateJWTCommand)
	refreshAuthCommand := auth.NewRefreshCommandHandler(logger, refreshTokenRepo, generateJWTCommand)

	getAccountCardsQuery := accounts.NewGetAccountCardsHandler(accountAPI)
	getAccountForProfileQuery := accounts.NewGetAccountForProfileQueryHandler(logger, accountAPI, customerRepo, accountRepo)
	getAccountListQuery := accounts.NewGetAccountListQueryHandler(logger, accountAPI)

	getCardNumberQuery := card.NewGetCardNumberQueryHandler(cardAPI, decryptService, logger)
	getCardCVVQuery := card.NewGetCardCVVQueryHandler(cardAPI, decryptService, logger)
	getCard3DSQuery := card.NewGetCard3DSQueryHandler(cardAPI, decryptService, logger)
	getCardPINQuery := card.NewGetCardPINQueryHandler(cardAPI, decryptService, logger)

	getCardTransactionsQueryHandler := transactions.NewGetCardTransactionsQueryHandler(transactionService)

	//handlers
	authHandler := auth2.NewAuthHandler(authCommand, registerCommand, refreshAuthCommand)
	accountHandler := acoount.NewAccountHandler(
		logger,
		getAccountCardsQuery,
		getAccountForProfileQuery,
		getAccountListQuery,
	)
	paymentHandler := payment.NewPaymentHandler(*paymentService)
	transactionHandler := transaction.NewTransactionHandler(getCardTransactionsQueryHandler)
	cardHandler := card2.NewCardHandler(
		logger,
		getCardNumberQuery,
		getCard3DSQuery,
		getCardPINQuery,
		getCardCVVQuery,
	)

	operationsHandler := operation.NewFinancialOperationHandler(logger)

	router := routes.NewRouter(
		logger,
		authHandler,
		accountHandler,
		paymentHandler,
		transactionHandler,
		cardHandler,
		operationsHandler,
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

func setupLogger(env string) *zerolog.Logger {
	var log zerolog.Logger

	logWriter := zerolog.MultiLevelWriter(os.Stdout)

	switch env {
	case envLocal:
		zerolog.SetGlobalLevel(zerolog.TraceLevel)
	case envDevelop:
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case envProduction:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	default:
		panic("unknown environment: " + env)
	}

	log = zerolog.New(logWriter).With().Caller().Logger()

	return &log
}
