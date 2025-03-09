package handlers

import (
	"decard/internal/application/query/transactions"
	"decard/internal/presentation/http/common"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type TransactionHandler struct {
	getCardTransactionsQueryHandler *transactions.GetCardTransactionsQueryHandler
}

func NewTransactionHandler(getCardTransactionsQueryHandler *transactions.GetCardTransactionsQueryHandler) *TransactionHandler {
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

	transactions, err := h.getCardTransactionsQueryHandler.Handle(transactions.GetCardTransactionsQuery{CardUUID: cardUUID})

	if err != nil {
		return err
	}

	return common.JSONResponse(w, http.StatusOK, transactions)
}
