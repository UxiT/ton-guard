package routes

import (
	"decard/internal/presentation/http/common"
	"decard/internal/presentation/http/handlers"
	"decard/internal/presentation/http/middleware"
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"
)

type APIFunc func(w http.ResponseWriter, r *http.Request) error

func Make(l *slog.Logger, h APIFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			e := common.JSONErrorResponse(w, err)

			if e != nil {
				http.Error(w, "internal server error", http.StatusInternalServerError)
			}

			l.Error(
				"http error",
				slog.String("method", r.Method),
				slog.String("url", r.URL.String()),
				slog.String("error", err.Error()),
			)
		}
	}
}

func NewRouter(
	logger *slog.Logger,
	authHandler *handlers.AuthHandler,
	accountHandler *handlers.AccountHandler,
	paymentHandler *handlers.PaymentHandler,
	transactionHandler *handlers.TransactionHandler,
	cardHandler *handlers.CardHandler,
) *mux.Router {
	r := mux.NewRouter()

	publicV1 := r.PathPrefix("/api/v1").Subrouter()

	// Public routes
	publicV1.HandleFunc("/auth/login", Make(logger, authHandler.Login)).Methods("POST")
	publicV1.HandleFunc("/auth/refresh", Make(logger, authHandler.Refresh)).Methods("POST")
	publicV1.HandleFunc("/auth/register", Make(logger, authHandler.Register)).Methods("POST")

	protectedV1 := r.PathPrefix("/api/v1").Subrouter()
	protectedV1.Use(middleware.AuthMiddleware)

	// Protected routes
	protectedV1.HandleFunc("/account", Make(logger, accountHandler.GetCustomerAccount)).Methods("GET")
	protectedV1.HandleFunc("/accounts", Make(logger, accountHandler.GetList)).Methods("GET")
	protectedV1.HandleFunc("/account/{account}/cards", Make(logger, accountHandler.GetAccountCards)).Methods("GET")

	protectedV1.HandleFunc("/transfer", Make(logger, paymentHandler.Transfer)).Methods("POST")

	protectedV1.HandleFunc("/transactions/{card}", Make(logger, transactionHandler.GetTransactionsByCard)).Methods("GET")

	//Card
	protectedV1.HandleFunc("/api/v1/cards", Make(logger, cardHandler.GetCustomerCards)).Methods("GET")
	protectedV1.HandleFunc("/api/v1/cards", Make(logger, cardHandler.Issue)).Methods("POST")
	protectedV1.HandleFunc("/api/v1/cards/{card}", Make(logger, cardHandler.Info)).Methods("GET")
	protectedV1.HandleFunc("/api/v1/cards/{card}/freeze", Make(logger, cardHandler.Freeze)).Methods("POST")
	protectedV1.HandleFunc("/api/v1/cards/{card}/block", Make(logger, cardHandler.Block)).Methods("POST")

	//Account
	//protected.HandleFunc("/api/v1/account", accountHandler.GetByCustomer).Methods("GET")

	//Card

	//Balance
	//protected.HandleFunc("/api/v1/balance", balanceHandler.GetByCustomer).Methods("GET")

	//Transactions

	return r
}
