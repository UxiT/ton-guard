package entity

import (
	"decard/internal/domain/valueobject"
	"time"
)

// Transaction represents a transaction record.
type Transaction struct {
	UUID      valueobject.UUID `json:"uuid"`
	CreatedAt time.Time        `json:"created_at"`
}
