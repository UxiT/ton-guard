package handlers

import (
	"decard/internal/domain"
	"decard/internal/domain/valueobject"
	"decard/internal/presentation/http/middleware"
	"github.com/rs/zerolog"
	"net/http"
)

type FinancialOperationHandler struct {
	logger *zerolog.Logger
}

func NewFinancialOperationHandler(logger *zerolog.Logger) *FinancialOperationHandler {
	return &FinancialOperationHandler{
		logger: logger,
	}
}

type topUpRequest struct {
	Amount  string `json:"amount"`
	Network string `json:"network"`
}

func (h FinancialOperationHandler) TopUp(w http.ResponseWriter, r *http.Request) error {
	const op = "http.handler.GetCustomerAccount"

	logger := h.logger.With().Str("operation", op).Logger()
	_, ok := r.Context().Value(middleware.ProfileUUIDKey).(valueobject.UUID)

	if !ok {
		logger.Error().Err(domain.ErrInvalidUser).Msg("failed to assert user UUID")

		return domain.ErrInvalidUser
	}

	return nil
}
