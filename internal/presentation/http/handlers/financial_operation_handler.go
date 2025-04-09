package handlers

import (
	"decard/internal/application/command/topup"
	"decard/internal/domain"
	"decard/internal/domain/valueobject"
	"decard/internal/presentation/http/middleware"
	"encoding/json"
	"github.com/rs/zerolog"
	"net/http"
)

type FinancialOperationHandler struct {
	logger             *zerolog.Logger
	createTopUpHandler topup.CreateCommandHandler
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

	var request topUpRequest

	logger := h.logger.With().Str("operation", op).Logger()
	_, ok := r.Context().Value(middleware.ProfileUUIDKey).(valueobject.UUID)

	if !ok {
		logger.Error().Err(domain.ErrInvalidUser).Msg("failed to assert user UUID")

		return domain.ErrInvalidUser
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return err
	}

	return h.createTopUpHandler.Handle(topup.CreateCommand{})
}
