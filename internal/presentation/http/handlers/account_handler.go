package handlers

import (
	"decard/internal/application/query"
	"decard/internal/domain/service"
	presentation "decard/internal/presentation/http"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type AccountHandler struct {
	accountService       service.AccountService
	getAccountCardsQuery query.GetAccountCardsHandler
}

func NewAccountHandler(
	accountService service.AccountService,
	getAccountCardsQuery query.GetAccountCardsHandler,
) *AccountHandler {
	return &AccountHandler{
		accountService:       accountService,
		getAccountCardsQuery: getAccountCardsQuery,
	}
}

func (h *AccountHandler) GetList(w http.ResponseWriter, r *http.Request) {
	accounts, err := h.accountService.GetAccountList()

	if err != nil {
		presentation.WriteError(w, err.Error(), http.StatusBadRequest)
	}

	presentation.WriteJSONResponse(w, accounts)
}

func (h *AccountHandler) GetAccountCards(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)

	accountUUID, ok := vars["account"]

	if !ok {
		return fmt.Errorf("invalid account UUID")
	}

	cards, err := h.getAccountCardsQuery.Handle(query.GetAccountCards{AccountUUID: accountUUID})
	if err != nil {
		return err
	}

	presentation.WriteJSONResponse(w, cards)

	return nil
}
