package interfaces

import (
	providerEntity "decard/internal/domain/entity/provider"
	"github.com/google/uuid"
)

type CardService interface {
	GetCustomerCards(account uuid.UUID) (*[]providerEntity.Card, error)
}

type AccountService interface {
	GetAccountsList() ([]providerEntity.Account, error)
	GetAccountCards(account uuid.UUID) ([]providerEntity.Card, error)
}

type TransactionService interface {
	GetCardTransactions(card uuid.UUID) (*[]providerEntity.Transaction, error)
}

type PaymentService interface {
	CreateAccountTransfer(amount float64, description, from, to string) error
}
