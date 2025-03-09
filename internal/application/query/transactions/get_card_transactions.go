package transactions

import (
	providerEntity "decard/internal/domain/entity/provider"
	"decard/internal/domain/service"
	"decard/internal/domain/valueobject"
)

type GetCardTransactionsQuery struct {
	CardUUID string
}

type GetCardTransactionsQueryHandler struct {
	transactionService *service.TransactionService
}

func NewGetCardTransactionsQueryHandler(transactionService *service.TransactionService) *GetCardTransactionsQueryHandler {
	return &GetCardTransactionsQueryHandler{
		transactionService: transactionService,
	}
}

func (h *GetCardTransactionsQueryHandler) Handle(q GetCardTransactionsQuery) (*[]providerEntity.Transaction, error) {
	cardUUID, err := valueobject.ParseUUID(q.CardUUID)

	if err != nil {
		return nil, err
	}

	return h.transactionService.GetTransactionsByCard(cardUUID)
}
