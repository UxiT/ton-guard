package handlers

import (
	"decard/internal/application/query/card"
	"decard/internal/domain"
	"decard/internal/domain/valueobject"
	"decard/internal/presentation/http/common"
	"decard/internal/presentation/http/middleware"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	"net/http"
)

type CardHandler struct {
	logger               *zerolog.Logger
	getCardNumberHandler *card.GetCardNumberQueryHandler
}

func NewCardHandler(logger *zerolog.Logger, getCardNumberHandler *card.GetCardNumberQueryHandler) *CardHandler {
	return &CardHandler{
		logger:               logger,
		getCardNumberHandler: getCardNumberHandler,
	}
}

func (h *CardHandler) GetCustomerCards(w http.ResponseWriter, r *http.Request) error {
	const op = "http.handler.GetCustomerCards"

	logger := h.logger.With().Str("operation", op).Logger()
	_, ok := r.Context().Value(middleware.ProfileUUIDKey).(valueobject.UUID)

	if !ok {
		logger.Error().Err(domain.ErrInvalidUser).Msg("failed to assert user UUID")

		return domain.ErrInvalidUser
	}

	return nil
}

func (h *CardHandler) Info(w http.ResponseWriter, r *http.Request) error {
	const op = "http.handler.cardInfo"

	logger := h.logger.With().Str("operation", op).Logger()
	_, ok := r.Context().Value(middleware.ProfileUUIDKey).(valueobject.UUID)

	if !ok {
		logger.Error().Err(domain.ErrInvalidUser).Msg("failed to assert user UUID")
		return domain.ErrInvalidUser
	}

	return nil
}

func (h *CardHandler) GetNumber(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	cardUUID, ok := vars["card"]

	if !ok {
		return fmt.Errorf("invalid card UUID")
	}

	number, err := h.getCardNumberHandler.Handle(card.GetCardNumberQuery{
		CardUUID: cardUUID,
	})

	if err != nil {
		return err
	}

	return common.JSONResponse(w, http.StatusOK, number)
}

func (h *CardHandler) Freeze(w http.ResponseWriter, r *http.Request) error {
	const op = "http.handler.cardFreeze"

	logger := h.logger.With().Str("operation", op).Logger()
	_, ok := r.Context().Value(middleware.ProfileUUIDKey).(valueobject.UUID)

	if !ok {
		logger.Error().Err(domain.ErrInvalidUser).Msg("failed to assert user UUID")
		return domain.ErrInvalidUser
	}

	return nil
}

func (h *CardHandler) Block(w http.ResponseWriter, r *http.Request) error {
	const op = "http.handler.cardBlock"

	logger := h.logger.With().Str("operation", op).Logger()
	_, ok := r.Context().Value(middleware.ProfileUUIDKey).(valueobject.UUID)

	if !ok {
		logger.Error().Err(domain.ErrInvalidUser).Msg("failed to assert user UUID")

		return domain.ErrInvalidUser
	}

	return nil
}

func (h *CardHandler) Issue(w http.ResponseWriter, r *http.Request) error {
	const op = "http.handler.cardIssue"

	logger := h.logger.With().Str("operation", op).Logger()
	_, ok := r.Context().Value(middleware.ProfileUUIDKey).(valueobject.UUID)

	if !ok {
		logger.Error().Err(domain.ErrInvalidUser).Msg("failed to assert user UUID")

		return domain.ErrInvalidUser
	}

	return nil
}
