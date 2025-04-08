package entity

import (
	"decard/internal/domain/valueobject"
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
	Profile   valueobject.UUID
	Amount    string
	Network   string
	Status    TopUpStatus
	IsClosed  bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
