package routes

import (
	"decard/internal/interfaces/http/handlers"
	"decard/internal/interfaces/http/middleware"

	"github.com/gorilla/mux"
)

func NewRouter(authHandler *handlers.AuthHandler) *mux.Router {
	r := mux.NewRouter()

	// Public routes
	r.HandleFunc("/api/v1/auth/login", authHandler.Login).Methods("POST")
	r.HandleFunc("/api/v1/auth/register", authHandler.Register).Methods("POST")

	// Protected routes
	protected := r.PathPrefix("/api").Subrouter()
	protected.Use(middleware.AuthMiddleware)

	//Account
	//protected.HandleFunc("/api/v1/account", accountHandler.GetByCustomer).Methods("GET")

	//Card

	//Balance
	//protected.HandleFunc("/api/v1/balance", balanceHandler.GetByCustomer).Methods("GET")

	//Transactions

	return r
}
