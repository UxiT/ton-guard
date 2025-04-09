package routes

import (
	"decard/internal/domain"
	"decard/internal/domain/valueobject"
	"decard/internal/presentation/http/common"
	"decard/internal/presentation/http/handlers"
	"decard/internal/presentation/http/handlers/acoount"
	"decard/internal/presentation/http/handlers/acoount/request"
	"decard/internal/presentation/http/middleware"
	"encoding/json"
	"github.com/rs/zerolog"
	"net/http"

	"github.com/gorilla/mux"
)

// Изменить так, чтобы можно было прокидывать структуру кастомного реквеста, он сам
// валидируется и прокидывается в функцию + настройка логгера тоже на этом уровне
type publicAPIFunc func(w http.ResponseWriter, r common.Request) error
type protectedAPIFunc func(w http.ResponseWriter, r common.Request, profile valueobject.UUID) error

func HandlePublic(l *zerolog.Logger, h publicAPIFunc, req common.Request) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			l.Error().Err(err).Msg("failed to parse request body")
			http.Error(w, "invalid request body", http.StatusUnprocessableEntity)

			return
		}

		if err := h(w, req); err != nil {
			e := common.JSONErrorResponse(w, err)

			if e != nil {
				http.Error(w, "internal server error", http.StatusInternalServerError)
			}

			l.Error().
				Str("method", r.Method).
				Str("url", r.URL.String()).
				Err(err).
				Msg("http error")
		}
	}
}

func HandleProtected(l *zerolog.Logger, h protectedAPIFunc, req common.Request) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := l.With().Str("method", r.Method).Str("url", r.URL.String()).Logger()
		profileUUID, ok := r.Context().Value(middleware.ProfileUUIDKey).(valueobject.UUID)

		if !ok {
			l.Error().Err(domain.ErrInvalidUser).Msg("failed to assert user UUID")

			e := common.JSONErrorResponse(w, domain.ErrUnauthorized)
			if e != nil {
				http.Error(w, "internal server error", http.StatusInternalServerError)
			}

			return
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			l.Error().Err(err).Msg("failed to parse request body")
			http.Error(w, "invalid request body", http.StatusUnprocessableEntity)

			return
		}

		if err := h(w, req, profileUUID); err != nil {
			e := common.JSONErrorResponse(w, err)

			if e != nil {
				http.Error(w, "internal server error", http.StatusInternalServerError)
			}

			l.Error().Err(err).Msg("http error")
		}
	}
}

func NewRouter(
	logger *zerolog.Logger,
	authHandler *handlers.AuthHandler,
	accountHandler *acoount.AccountHandler,
	paymentHandler *handlers.PaymentHandler,
	transactionHandler *handlers.TransactionHandler,
	cardHandler *handlers.CardHandler,
	operationsHandler *handlers.FinancialOperationHandler,
) *mux.Router {
	r := mux.NewRouter()

	publicV1 := r.PathPrefix("/api/v1").Subrouter()

	// Public routes
	publicV1.HandleFunc("/auth/login", HandlePublic(logger, authHandler.Login)).Methods("POST")
	publicV1.HandleFunc("/auth/refresh", HandlePublic(logger, authHandler.Refresh)).Methods("POST")
	publicV1.HandleFunc("/auth/register", HandlePublic(logger, authHandler.Register)).Methods("POST")

	protectedV1 := r.PathPrefix("/api/v1").Subrouter()
	protectedV1.Use(middleware.AuthMiddleware)

	// Protected routes
	protectedV1.HandleFunc("/account", HandleProtected(logger, accountHandler.GetCustomerAccount, nil)).Methods("GET")
	protectedV1.HandleFunc("/accounts", HandleProtected(logger, accountHandler.GetList, nil)).Methods("GET")
	protectedV1.HandleFunc("/account/{account}/cards", HandleProtected(logger, accountHandler.GetAccountCards, request.GetAccountCardsRequest{})).Methods("GET")

	protectedV1.HandleFunc("/transfer", HandleProtected(logger, paymentHandler.Transfer)).Methods("POST")

	protectedV1.HandleFunc("/transactions/{card}", HandleProtected(logger, transactionHandler.GetTransactionsByCard)).Methods("GET")

	//Card
	protectedV1.HandleFunc("/cards", HandleProtected(logger, cardHandler.GetCustomerCards)).Methods("GET")
	protectedV1.HandleFunc("/cards", HandleProtected(logger, cardHandler.Issue)).Methods("POST")
	protectedV1.HandleFunc("/cards/{card}", HandleProtected(logger, cardHandler.Info)).Methods("GET")

	protectedV1.HandleFunc("/cards/{card}/number", HandleProtected(logger, cardHandler.GetNumber)).Methods("GET")
	protectedV1.HandleFunc("/cards/{card}/3ds", HandleProtected(logger, cardHandler.Get3DS)).Methods("GET")
	protectedV1.HandleFunc("/cards/{card}/cvv", HandleProtected(logger, cardHandler.GetCVV)).Methods("GET")
	protectedV1.HandleFunc("/cards/{card}/pin", HandleProtected(logger, cardHandler.GetPIN)).Methods("GET")

	protectedV1.HandleFunc("/cards/{card}/freeze", HandleProtected(logger, cardHandler.Freeze)).Methods("POST")
	protectedV1.HandleFunc("/cards/{card}/block", HandleProtected(logger, cardHandler.Block)).Methods("POST")

	//TopUp
	protectedV1.HandleFunc("/operations/top-up", HandleProtected(logger, operationsHandler.TopUp)).Methods("POST")
	protectedV1.HandleFunc("/operations/top-up/cancel", HandleProtected(logger, operationsHandler.TopUp)).Methods("GET")
	protectedV1.HandleFunc("/operations/top-up/decline", HandleProtected(logger, operationsHandler.TopUp)).Methods("GET")
	protectedV1.HandleFunc("/operations/top-up/add_trx_id", HandleProtected(logger, operationsHandler.TopUp)).Methods("POST")
	protectedV1.HandleFunc("/operations/top-up/approve", HandleProtected(logger, operationsHandler.TopUp)).Methods("GET")
	protectedV1.HandleFunc("/operations/top-up/close", HandleProtected(logger, operationsHandler.TopUp)).Methods("GET")

	//Account
	//protected.HandleFunc("/api/v1/account", accountHandler.GetByCustomer).Methods("GET")

	//Card

	//Balance
	//protected.HandleFunc("/api/v1/balance", balanceHandler.GetByCustomer).Methods("GET")

	//Transactions

	return r
}
