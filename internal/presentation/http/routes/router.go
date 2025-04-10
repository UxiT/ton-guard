package routes

import (
	"decard/internal/presentation/http/handlers"
	"decard/internal/presentation/http/handlers/acoount"
	"decard/internal/presentation/http/handlers/auth"
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
	cardHandler *handlers.CardHandler,
	operationsHandler *handlers.FinancialOperationHandler,
) *mux.Router {
	r := mux.NewRouter()

	publicV1 := r.PathPrefix("/api/v1").Subrouter()

	// Public routes
	publicV1.HandleFunc("/auth/login", HandlePublic(logger, authHandler.Login, auth.LoginRequest{})).Methods("POST")
	publicV1.HandleFunc("/auth/refresh", HandlePublic(logger, authHandler.Refresh, auth.RefreshRequest{})).Methods("POST")
	publicV1.HandleFunc("/auth/register", HandlePublic(logger, authHandler.Register, auth.RegisterRequest{})).Methods("POST")

	protectedV1 := r.PathPrefix("/api/v1").Subrouter()
	protectedV1.Use(middleware.AuthMiddleware)

	// Protected routes
	protectedV1.HandleFunc("/account", HandleProtected(logger, accountHandler.GetCustomerAccount, nil)).Methods("GET")
	protectedV1.HandleFunc("/accounts", HandleProtected(logger, accountHandler.GetList, nil)).Methods("GET")
	protectedV1.HandleFunc("/account/{account}/cards", HandleProtected(logger, accountHandler.GetAccountCards, acoount.GetAccountCardsRequest{})).Methods("GET")

	protectedV1.HandleFunc("/transfer", HandleProtected(logger, paymentHandler.Transfer, payment.TransferRequest{})).Methods("POST")

	protectedV1.HandleFunc("/transactions/{card}", HandleProtected(logger, transactionHandler.GetTransactionsByCard, transaction.GetCardTransactionRequest{})).Methods("GET")

	//Card
	//protectedV1.HandleFunc("/cards", HandleProtected(logger, cardHandler.GetCustomerCards)).Methods("GET")
	//protectedV1.HandleFunc("/cards", HandleProtected(logger, cardHandler.Issue)).Methods("POST")
	//protectedV1.HandleFunc("/cards/{card}", HandleProtected(logger, cardHandler.Info)).Methods("GET")
	//
	//protectedV1.HandleFunc("/cards/{card}/number", HandleProtected(logger, cardHandler.GetNumber)).Methods("GET")
	//protectedV1.HandleFunc("/cards/{card}/3ds", HandleProtected(logger, cardHandler.Get3DS)).Methods("GET")
	//protectedV1.HandleFunc("/cards/{card}/cvv", HandleProtected(logger, cardHandler.GetCVV)).Methods("GET")
	//protectedV1.HandleFunc("/cards/{card}/pin", HandleProtected(logger, cardHandler.GetPIN)).Methods("GET")
	//
	//protectedV1.HandleFunc("/cards/{card}/freeze", HandleProtected(logger, cardHandler.Freeze)).Methods("POST")
	//protectedV1.HandleFunc("/cards/{card}/block", HandleProtected(logger, cardHandler.Block)).Methods("POST")
	//
	////TopUp
	//protectedV1.HandleFunc("/operations/top-up", HandleProtected(logger, operationsHandler.TopUp)).Methods("POST")
	//protectedV1.HandleFunc("/operations/top-up/cancel", HandleProtected(logger, operationsHandler.TopUp)).Methods("GET")
	//protectedV1.HandleFunc("/operations/top-up/decline", HandleProtected(logger, operationsHandler.TopUp)).Methods("GET")
	//protectedV1.HandleFunc("/operations/top-up/add_trx_id", HandleProtected(logger, operationsHandler.TopUp)).Methods("POST")
	//protectedV1.HandleFunc("/operations/top-up/approve", HandleProtected(logger, operationsHandler.TopUp)).Methods("GET")
	//protectedV1.HandleFunc("/operations/top-up/close", HandleProtected(logger, operationsHandler.TopUp)).Methods("GET")

	//Account
	//protected.HandleFunc("/api/v1/account", accountHandler.GetByCustomer).Methods("GET")

	//Card

	//Balance
	//protected.HandleFunc("/api/v1/balance", balanceHandler.GetByCustomer).Methods("GET")

	//Transactions

	return r
}
