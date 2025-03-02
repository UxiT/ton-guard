package routes

import (
	http2 "decard/internal/presentation/http"
	"decard/internal/presentation/http/handlers"
	"decard/internal/presentation/http/middleware"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type APIError struct {
	StatusCode int `json:"status_code"`
	Msg        any `json:"msg"`
}

func (e APIError) Error() string {
	return fmt.Sprintf("api error: %d", e.StatusCode)
}

func NewApiError(statusCode int, err error) APIError {
	return APIError{
		StatusCode: statusCode,
		Msg:        err.Error(),
	}
}

type APIFunc func(w http.ResponseWriter, r *http.Request) error

func Make(h APIFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			if apiErr, ok := err.(APIError); ok {
				http2.WriteError(w, apiErr.Error(), apiErr.StatusCode)
			} else {
				errResp := map[string]any{
					"status_code": http.StatusInternalServerError,
					"msg":         "internal server error",
				}

				log.Printf("error: %s", err.Error())
				http2.WriteError(w, errResp, http.StatusInternalServerError)
			}
		}
	}
}

func NewRouter(
	authHandler *handlers.AuthHandler,
	accountHandler *handlers.AccountHandler,
	paymentHandler *handlers.PaymentHandler,
	transactionHandler *handlers.TransactionHandler,
) *mux.Router {
	r := mux.NewRouter()

	// Public routes
	r.HandleFunc("/api/v1/auth/login", authHandler.Login).Methods("POST")
	r.HandleFunc("/api/v1/auth/register", authHandler.Register).Methods("POST")

	r.HandleFunc("/api/v1/accounts", accountHandler.GetList).Methods("GET")
	r.HandleFunc("/api/v1/account/{account}/cards", Make(accountHandler.GetAccountCards)).Methods("GET")

	r.HandleFunc("/api/v1/transfer", paymentHandler.Transfer).Methods("POST")

	r.HandleFunc("/api/v1/transactions/{card}", Make(transactionHandler.GetTransactionsByCard)).Methods("GET")
	// Protected routes
	protected := r.PathPrefix("/api").Subrouter()
	protected.Use(middleware.AuthMiddleware)

	//Card
	//protected.HandleFunc("/api/v1/cards", cardHandler.GetCardList).Methods("GET")

	//Account
	//protected.HandleFunc("/api/v1/account", accountHandler.GetByCustomer).Methods("GET")

	//Card

	//Balance
	//protected.HandleFunc("/api/v1/balance", balanceHandler.GetByCustomer).Methods("GET")

	//Transactions

	return r
}
