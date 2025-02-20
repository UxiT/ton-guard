package entity

import (
	"time"

	"github.com/google/uuid"
)

// Transaction represents a transaction record.
type Transaction struct {
	UUID      uuid.UUID `json:"uuid"`
	CreatedAt time.Time `json:"created_at"`
}
