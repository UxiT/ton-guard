package card

import (
	"decard/internal/application/query/card"
	"decard/internal/domain/valueobject"
	"decard/internal/presentation/http/common"
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

func (h *CardHandler) Info(w http.ResponseWriter, r any, profileUUID valueobject.UUID) error {
	const op = "http.handler.cardInfo"

	return nil
}

func (h *CardHandler) GetNumber(w http.ResponseWriter, r any, profile valueobject.UUID) error {
	req := r.(GetCardInfoRequest)

	number, err := h.getCardNumberHandler.Handle(card.GetCardNumberQuery{
		CardUUID: req.Card,
	})

	if err != nil {
		return err
	}

	return common.JSONResponse(w, http.StatusOK, number)
}

func (h *CardHandler) Get3DS(w http.ResponseWriter, r any, profile valueobject.UUID) error {
	req := r.(GetCardInfoRequest)

	number, err := h.getCard3DSHandler.Handle(card.GetCard3DSQuery{
		CardUUID: req.Card,
	})

	if err != nil {
		return err
	}

	return common.JSONResponse(w, http.StatusOK, number)
}

func (h *CardHandler) GetCVV(w http.ResponseWriter, r any, profile valueobject.UUID) error {
	req := r.(GetCardInfoRequest)

	number, err := h.getCardCVVHandler.Handle(card.GetCardCVVQuery{
		CardUUID: req.Card,
	})

	if err != nil {
		return err
	}

	return common.JSONResponse(w, http.StatusOK, number)
}

func (h *CardHandler) GetPIN(w http.ResponseWriter, r any, profile valueobject.UUID) error {
	req := r.(GetCardInfoRequest)

	number, err := h.getCardPINHandler.Handle(card.GetCardPINQuery{
		CardUUID: req.Card,
	})

	if err != nil {
		return err
	}

	return common.JSONResponse(w, http.StatusOK, number)
}

func (h *CardHandler) Freeze(w http.ResponseWriter, r any, profile valueobject.UUID) error {
	const op = "http.handler.cardFreeze"

	return nil
}

func (h *CardHandler) Block(w http.ResponseWriter, r any, profile valueobject.UUID) error {
	const op = "http.handler.cardBlock"

	return nil
}

func (h *CardHandler) Issue(w http.ResponseWriter, r any, profile valueobject.UUID) error {
	const op = "http.handler.cardIssue"

	return nil
}
