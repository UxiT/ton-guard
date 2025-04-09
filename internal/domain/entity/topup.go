package entity

import (
	"decard/internal/domain/valueobject"
	"github.com/shopspring/decimal"
	"time"
)

type TopUpStatus string

const (
	WaitingForTxID   TopUpStatus = "waiting_for_tx_id"
	Cancelled        TopUpStatus = "cancelled"
	Validating       TopUpStatus = "validating"
	ValidationFailed TopUpStatus = "validation_failed"
	TopUpInProgress  TopUpStatus = "top_up_in_progress"
	Completed        TopUpStatus = "completed"
)

type TopUp struct {
	UUID      valueobject.UUID
	Customer  valueobject.UUID
	Amount    decimal.Decimal
	Network   string
	Status    TopUpStatus
	IsClosed  bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewTopUp(customer valueobject.UUID, amount decimal.Decimal, network string) TopUp {
	return TopUp{
		UUID:      valueobject.NewUUID(),
		Customer:  customer,
		Amount:    amount,
		Network:   network,
		Status:    WaitingForTxID,
		IsClosed:  false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
