package handlers

import (
	"decard/internal/domain"
	"decard/internal/domain/valueobject"
	"decard/internal/presentation/http/middleware"
	"log/slog"
	"net/http"
)

type CardHandler struct {
	logger *slog.Logger
}

func NewCardHandler(logger *slog.Logger) *CardHandler {
	return &CardHandler{
		logger: logger,
	}
}

func (h *CardHandler) GetCustomerCards(w http.ResponseWriter, r *http.Request) error {
	const op = "http.handler.GetCustomerCards"

	logger := h.logger.With(slog.String("operation", op))
	_, ok := r.Context().Value(middleware.ProfileUUIDKey).(valueobject.UUID)

	if !ok {
		logger.Error("failed to assert user UUID")
		return domain.ErrInvalidUser
	}

	return nil
}
