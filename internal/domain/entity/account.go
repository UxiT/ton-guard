package entity

import (
	domaintype "decard/internal/domain/type"
	"decard/internal/domain/valueobject"
	"fmt"
	"time"

	"github.com/shopspring/decimal"
)

type Account struct {
	UUID         valueobject.UUID
	ExternalUUID valueobject.UUID
	Currency     domaintype.Currency
	Status       domaintype.AccountStatus
	Balance      Balance
	CustomerUUID valueobject.UUID

	CreatedAt time.Time
	UpdatedAt time.Time
}

type Balance decimal.Decimal

func CreateAccount(
	externalUUID valueobject.UUID,
	balance float64,
	currency string,
) (*Account, error) {
	b, err := NewBalance(fmt.Sprintf("%f", balance))

	if err != nil {
		return nil, err
	}

	return &Account{
		UUID:         valueobject.NewUUID(),
		ExternalUUID: externalUUID,
		Currency:     domaintype.Currency(currency),
		Status:       domaintype.AccountActive,
		Balance:      *b,
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
