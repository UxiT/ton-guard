package entity

import (
	"decard/internal/domain/valueobject"
	"time"

	"github.com/shopspring/decimal"
)

type Account struct {
	UUID         valueobject.UUID
	ExternalUUID valueobject.UUID
	Currency     Currency
	Status       AccountStatus
	Balance      Balance
	CardUUID     valueobject.UUID

	CreatedAt time.Time
	UpdatedAt time.Time
}

type Currency string
type AccountStatus string

type Balance decimal.Decimal

func CreateAccount(
	externalUUID valueobject.UUID,
	currency Currency,
	status AccountStatus,
) (*Account, error) {
	balance, err := NewBalance("0")

	if err != nil {
		return nil, err
	}

	return &Account{
		UUID:         valueobject.NewUUID(),
		ExternalUUID: externalUUID,
		Currency:     currency,
		Status:       status,
		Balance:      *balance,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}, nil
}

func NewBalance(i string) (*Balance, error) {
	balance, err := decimal.NewFromString(i)
	if err != nil {
		return nil, err
	}

	return (*Balance)(&balance), nil
}
