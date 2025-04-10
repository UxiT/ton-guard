package transaction

import (
	"decard/internal/application/query/transactions"
	"decard/internal/domain/valueobject"
	"decard/internal/presentation/http/common"
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

func (h *TransactionHandler) GetTransactionsByCard(w http.ResponseWriter, r any, profileUUID valueobject.UUID) error {
	req := r.(*GetCardTransactionRequest)

	trns, err := h.getCardTransactionsQueryHandler.Handle(transactions.GetCardTransactionsQuery{CardUUID: req.Card})

	if err != nil {
		return err
	}

	return common.JSONResponse(w, http.StatusOK, trns)
}
