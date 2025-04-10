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
	getCard3DSHandler    *card.GetCard3DSQueryHandler
	getCardPINHandler    *card.GetCardPINQueryHandler
	getCardCVVHandler    *card.GetCardCVVQueryHandler
}

func NewCardHandler(
	logger *zerolog.Logger,
	getCardNumberHandler *card.GetCardNumberQueryHandler,
	getCard3DSHandler *card.GetCard3DSQueryHandler,
	getCardPINHandler *card.GetCardPINQueryHandler,
	getCardCVVHandler *card.GetCardCVVQueryHandler,
) *CardHandler {
	return &CardHandler{
		logger:               logger,
		getCardNumberHandler: getCardNumberHandler,
		getCard3DSHandler:    getCard3DSHandler,
		getCardPINHandler:    getCardPINHandler,
		getCardCVVHandler:    getCardCVVHandler,
	}
}

func (h *CardHandler) GetCustomerCards(w http.ResponseWriter, r any, profileUUID valueobject.UUID) error {
	const op = "http.handler.GetCustomerCards"

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

func (h *CardHandler) Get3DS(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	cardUUID, ok := vars["card"]

	if !ok {
		return fmt.Errorf("invalid card UUID")
	}

	number, err := h.getCard3DSHandler.Handle(card.GetCard3DSQuery{
		CardUUID: cardUUID,
	})

	if err != nil {
		return err
	}

	return common.JSONResponse(w, http.StatusOK, number)
}

func (h *CardHandler) GetCVV(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	cardUUID, ok := vars["card"]

	if !ok {
		return fmt.Errorf("invalid card UUID")
	}

	number, err := h.getCardCVVHandler.Handle(card.GetCardCVVQuery{
		CardUUID: cardUUID,
	})

	if err != nil {
		return err
	}

	return common.JSONResponse(w, http.StatusOK, number)
}

func (h *CardHandler) GetPIN(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	cardUUID, ok := vars["card"]

	if !ok {
		return fmt.Errorf("invalid card UUID")
	}

	number, err := h.getCardPINHandler.Handle(card.GetCardPINQuery{
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
