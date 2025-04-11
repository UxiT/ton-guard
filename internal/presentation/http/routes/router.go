package routes

import (
	"decard/internal/presentation/http/handlers/acoount"
	"decard/internal/presentation/http/handlers/auth"
	"decard/internal/presentation/http/handlers/card"
	"decard/internal/presentation/http/handlers/operation"
	"decard/internal/presentation/http/handlers/payment"
	"decard/internal/presentation/http/handlers/transaction"
	"decard/internal/presentation/http/middleware"
	"github.com/rs/zerolog"

	"github.com/gorilla/mux"
)

func NewRouter(
	logger *zerolog.Logger,
	authHandler *auth.AuthHandler,
	accountHandler *acoount.AccountHandler,
	paymentHandler *payment.PaymentHandler,
	transactionHandler *transaction.TransactionHandler,
	cardHandler *card.CardHandler,
	operationsHandler *operation.FinancialOperationHandler,
) *mux.Router {
	r := mux.NewRouter()

	publicV1 := r.PathPrefix("/api/v1").Subrouter()

	// Public routes
	publicV1.HandleFunc("/auth/login", HandlePublic(logger, authHandler.Login, auth.LoginRequest{})).Methods("POST")
	publicV1.HandleFunc("/auth/refresh", HandlePublic(logger, authHandler.Refresh, auth.RefreshRequest{})).Methods("POST")
	publicV1.HandleFunc("/auth/register", HandlePublic(logger, authHandler.Register, auth.RegisterRequest{})).Methods("POST")

	protectedV1 := r.PathPrefix("/api/v1").Subrouter()
	protectedV1.Use(middleware.AuthMiddleware)

	private := r.PathPrefix("/api/v1").Subrouter()

	// Protected routes
	protectedV1.HandleFunc("/account", HandleProtected(logger, accountHandler.GetCustomerAccount, nil)).Methods("GET")
	protectedV1.HandleFunc("/accounts", HandleProtected(logger, accountHandler.GetList, nil)).Methods("GET")
	protectedV1.HandleFunc("/account/{account}/cards", HandleProtected(logger, accountHandler.GetAccountCards, acoount.GetAccountCardsRequest{})).Methods("GET")

	protectedV1.HandleFunc("/transfer", HandleProtected(logger, paymentHandler.Transfer, payment.TransferRequest{})).Methods("POST")

	protectedV1.HandleFunc("/transactions/{card}", HandleProtected(logger, transactionHandler.GetTransactionsByCard, transaction.GetCardTransactionRequest{})).Methods("GET")

	//Card
	protectedV1.HandleFunc("/cards", HandleProtected(logger, cardHandler.GetCustomerCards, nil)).Methods("GET")
	protectedV1.HandleFunc("/cards", HandleProtected(logger, cardHandler.Issue, nil)).Methods("POST")

	protectedV1.HandleFunc("/cards/{card}", HandleProtected(logger, cardHandler.Info, card.GetCardInfoRequest{})).Methods("GET")
	protectedV1.HandleFunc("/cards/{card}/number", HandleProtected(logger, cardHandler.GetNumber, card.GetCardInfoRequest{})).Methods("GET")
	protectedV1.HandleFunc("/cards/{card}/3ds", HandleProtected(logger, cardHandler.Get3DS, card.GetCardInfoRequest{})).Methods("GET")
	protectedV1.HandleFunc("/cards/{card}/cvv", HandleProtected(logger, cardHandler.GetCVV, card.GetCardInfoRequest{})).Methods("GET")
	protectedV1.HandleFunc("/cards/{card}/pin", HandleProtected(logger, cardHandler.GetPIN, card.GetCardInfoRequest{})).Methods("GET")

	protectedV1.HandleFunc("/cards/{card}/freeze", HandleProtected(logger, cardHandler.Freeze, card.GetCardInfoRequest{})).Methods("POST")
	protectedV1.HandleFunc("/cards/{card}/block", HandleProtected(logger, cardHandler.Block, card.GetCardInfoRequest{})).Methods("POST")

	//TopUp
	protectedV1.HandleFunc("/operations/top-up", HandleProtected(logger, operationsHandler.CreateTopUp, operation.TopUpRequest{})).Methods("POST")
	protectedV1.HandleFunc("/operations/top-up/{uuid}/cancel", HandleProtected(logger, operationsHandler.Cancel, operation.TopUpUUIDRequest{})).Methods("GET")
	protectedV1.HandleFunc("/operations/top-up/{uuid}/add_trx_id", HandleProtected(logger, operationsHandler.AddTransactionID, operation.AddTransactionRequest{})).Methods("POST")

	private.HandleFunc("/operations/top-up/{uuid}/decline", HandlePrivate(logger, operationsHandler.Decline, operation.TopUpUUIDRequest{})).Methods("GET")
	private.HandleFunc("/operations/top-up/{uuid}/approve", HandlePrivate(logger, operationsHandler.Approve, operation.TopUpUUIDRequest{})).Methods("GET")
	private.HandleFunc("/operations/top-up/{uuid}/close", HandlePrivate(logger, operationsHandler.Close, operation.TopUpUUIDRequest{})).Methods("GET")

	//Account
	//protected.HandleFunc("/api/v1/account", accountHandler.GetByCustomer).Methods("GET")

	//Card

	//Balance
	//protected.HandleFunc("/api/v1/balance", balanceHandler.GetByCustomer).Methods("GET")

	//Transactions

	return r
}
