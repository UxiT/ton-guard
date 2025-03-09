package interfaces

import (
	providerEntity "decard/internal/domain/entity/provider"
	"decard/internal/domain/valueobject"
)

type CardService interface {
	GetCustomerCards(account valueobject.UUID) (*[]providerEntity.Card, error)
}

type AccountService interface {
	GetAccountsList() ([]providerEntity.Account, error)
	GetAccount(account valueobject.UUID) (*providerEntity.Account, error)
	GetAccountCards(account valueobject.UUID) ([]providerEntity.Card, error)
}

type TransactionService interface {
	GetCardTransactions(card valueobject.UUID) (*[]providerEntity.Transaction, error)
}

type PaymentService interface {
	CreateAccountTransfer(amount float64, description, from, to string) error
}
