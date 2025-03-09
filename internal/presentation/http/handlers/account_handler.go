package handlers

import (
	"decard/internal/application/query/accounts"
	"decard/internal/domain"
	"decard/internal/domain/valueobject"
	"decard/internal/presentation/http/common"
	"decard/internal/presentation/http/middleware"
	"fmt"
	"github.com/gorilla/mux"
	"log/slog"
	"net/http"
)

type AccountHandler struct {
	logger                    *slog.Logger
	getAccountCardsQuery      accounts.GetAccountCardsHandler
	getAccountForProfileQuery accounts.GetAccountForProfileQueryHandler
	getAccountListQuery       accounts.GetAccountListQueryHandler
}

func NewAccountHandler(
	logger *slog.Logger,
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

func (h *AccountHandler) GetCustomerAccount(w http.ResponseWriter, r *http.Request) error {
	const op = "http.handler.GetCustomerAccount"

	logger := h.logger.With(slog.String("operation", op))
	profileUUID, ok := r.Context().Value(middleware.ProfileUUIDKey).(valueobject.UUID)

	if !ok {
		logger.Error("failed to assert user UUID")
		return domain.ErrInvalidUser
	}

	account, err := h.getAccountForProfileQuery.Handle(accounts.GetAccountForProfileQuery{ProfileUUID: profileUUID})

	if err != nil {
		logger.Error("error getting customer account", slog.String("error", err.Error()))
		return err
	}

	return common.JSONResponse(w, http.StatusOK, account)

}

func (h *AccountHandler) GetList(w http.ResponseWriter, r *http.Request) error {
	accountList, err := h.getAccountListQuery.Handle()

	if err != nil {
		return err
	}

	return common.JSONResponse(w, http.StatusOK, accountList)
}

func (h *AccountHandler) GetAccountCards(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)

	accountUUID, ok := vars["account"]

	if !ok {
		return fmt.Errorf("invalid account UUID")
	}

	cards, err := h.getAccountCardsQuery.Handle(accounts.GetAccountCards{AccountUUID: accountUUID})
	if err != nil {
		return err
	}

	return common.JSONResponse(w, http.StatusOK, cards)
}
