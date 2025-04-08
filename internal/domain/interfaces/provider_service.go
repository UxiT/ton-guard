package interfaces

import (
	providerEntity "decard/internal/domain/entity/provider"
	"decard/internal/domain/valueobject"
)

type CardService interface {
	GetCard(card valueobject.UUID) (*providerEntity.Card, error)
	BlockCard(card valueobject.UUID) (*providerEntity.Card, error)
	CloseCard(card valueobject.UUID) (*providerEntity.Card, error)
	CreateCard() (*providerEntity.Card, error)

	GetCardNumber(card valueobject.UUID) (string, error)
	GetCardCVV(card valueobject.UUID) (string, error)
	GetCard3DS(card valueobject.UUID) (string, error)
	GetCardPIN(card valueobject.UUID) (string, error)
}

type AccountService interface {
	GetAccountsList() ([]providerEntity.Account, error)
	GetAccount(account valueobject.UUID) (*providerEntity.Account, error)
	GetAccountCards(account valueobject.UUID) ([]providerEntity.Card, error)
	CreateAccount(name string) (*providerEntity.Account, error)
}

type TransactionService interface {
	GetCardTransactions(card valueobject.UUID) (*[]providerEntity.Transaction, error)
}

type PaymentService interface {
	CreateAccountTransfer(amount float64, description, from, to string) error
}
