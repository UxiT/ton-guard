package service

import (
	providerEntity "decard/internal/domain/entity/provider"
	"decard/internal/domain/interfaces"
	"decard/internal/domain/valueobject"
)

type TransactionService struct {
	providerTransactionService interfaces.TransactionService
}

func NewTransactionService(providerTransactionService interfaces.TransactionService) *TransactionService {
	return &TransactionService{
		providerTransactionService: providerTransactionService,
	}
}

func (s *TransactionService) GetTransactionsByCard(card valueobject.UUID) (*[]providerEntity.Transaction, error) {
	//validate user has access to this card

	return s.providerTransactionService.GetCardTransactions(card)
}
