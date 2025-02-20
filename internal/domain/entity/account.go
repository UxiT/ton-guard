package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Account struct {
	UUID         uuid.UUID
	ExternalUUID uuid.UUID
	Currency     Currency
	Status       AccountStatus
	Balance      Balance

	CreatedAt time.Time
	UpdatedAt time.Time
}

type Currency string
type AccountStatus string

type Balance decimal.Decimal

func NewBalance(i string) (*Balance, error) {
	balance, err := decimal.NewFromString(i)
	if err != nil {
		return nil, err
	}

	return (*Balance)(&balance), nil
}
