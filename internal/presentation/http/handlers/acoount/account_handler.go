package acoount

import (
	"decard/internal/application/query/accounts"
	"decard/internal/domain/valueobject"
	"decard/internal/presentation/http/common"
	"github.com/rs/zerolog"
	"net/http"
)

type AccountHandler struct {
	logger                    *zerolog.Logger
	getAccountCardsQuery      accounts.GetAccountCardsHandler
	getAccountForProfileQuery accounts.GetAccountForProfileQueryHandler
	getAccountListQuery       accounts.GetAccountListQueryHandler
}

func NewAccountHandler(
	logger *zerolog.Logger,
	getAccountCardsQuery accounts.GetAccountCardsHandler,
	getAccountForProfileQuery accounts.GetAccountForProfileQueryHandler,
	getAccountListQuery accounts.GetAccountListQueryHandler,
) *AccountHandler {
	return &AccountHandler{
		logger:                    logger,
		getAccountCardsQuery:      getAccountCardsQuery,
		getAccountForProfileQuery: getAccountForProfileQuery,
		getAccountListQuery:       getAccountListQuery,
	}
}

func (h *AccountHandler) GetCustomerAccount(w http.ResponseWriter, r any, profileUUID valueobject.UUID) error {
	const op = "http.handler.GetCustomerAccount"

	logger := h.logger.With().Str("operation", op).Logger()

	account, err := h.getAccountForProfileQuery.Handle(accounts.GetAccountForProfileQuery{ProfileUUID: profileUUID})

	if err != nil {
		logger.Error().Err(err).Msg("error getting customer account")

		return err
	}

	return common.JSONResponse(w, http.StatusOK, account)

}

func (h *AccountHandler) GetList(w http.ResponseWriter, r any, profileUUID valueobject.UUID) error {
	accountList, err := h.getAccountListQuery.Handle()
	if err != nil {
		return err
	}

	return common.JSONResponse(w, http.StatusOK, accountList)
}

func (h *AccountHandler) GetAccountCards(w http.ResponseWriter, r any, profileUUID valueobject.UUID) error {
	req := r.(*GetAccountCardsRequest)

	cards, err := h.getAccountCardsQuery.Handle(accounts.GetAccountCards{AccountUUID: req.Account})
	if err != nil {
		return err
	}

	return common.JSONResponse(w, http.StatusOK, cards)
}
