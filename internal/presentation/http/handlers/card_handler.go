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

func (h *CardHandler) Info(w http.ResponseWriter, r *http.Request) error {
	const op = "http.handler.cardInfo"

	logger := h.logger.With(slog.String("operation", op))
	_, ok := r.Context().Value(middleware.ProfileUUIDKey).(valueobject.UUID)

	if !ok {
		logger.Error("failed to assert user UUID")
		return domain.ErrInvalidUser
	}

	return nil
}

func (h *CardHandler) Freeze(w http.ResponseWriter, r *http.Request) error {
	const op = "http.handler.cardFreeze"

	logger := h.logger.With(slog.String("operation", op))
	_, ok := r.Context().Value(middleware.ProfileUUIDKey).(valueobject.UUID)

	if !ok {
		logger.Error("failed to assert user UUID")
		return domain.ErrInvalidUser
	}

	return nil
}

func (h *CardHandler) Block(w http.ResponseWriter, r *http.Request) error {
	const op = "http.handler.cardBlock"

	logger := h.logger.With(slog.String("operation", op))
	_, ok := r.Context().Value(middleware.ProfileUUIDKey).(valueobject.UUID)

	if !ok {
		logger.Error("failed to assert user UUID")
		return domain.ErrInvalidUser
	}

	return nil
}

func (h *CardHandler) Issue(w http.ResponseWriter, r *http.Request) error {
	const op = "http.handler.cardIssue"

	logger := h.logger.With(slog.String("operation", op))
	_, ok := r.Context().Value(middleware.ProfileUUIDKey).(valueobject.UUID)

	if !ok {
		logger.Error("failed to assert user UUID")
		return domain.ErrInvalidUser
	}

	return nil
}
