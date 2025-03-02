package handlers

import (
	"decard/internal/application/query"
	presentation "decard/internal/presentation/http"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type TransactionHandler struct {
	getCardTransactionsQueryHandler *query.GetCardTransactionsQueryHandler
}

func NewTransactionHandler(getCardTransactionsQueryHandler *query.GetCardTransactionsQueryHandler) *TransactionHandler {
	return &TransactionHandler{
		getCardTransactionsQueryHandler: getCardTransactionsQueryHandler,
	}
}

func (h *TransactionHandler) GetTransactionsByCard(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	cardUUID, ok := vars["card"]

	if !ok {
		return fmt.Errorf("invalid card UUID")
	}

	transactions, err := h.getCardTransactionsQueryHandler.Handle(query.GetCardTransactionsQuery{CardUUID: cardUUID})

	if err != nil {
		return err
	}

	presentation.WriteJSONResponse(w, transactions)

	return nil
}
