package entity

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

// Batch represents a batch record.
type Batch struct {
	UUID        uuid.UUID       `json:"uuid"`
	FirstID     uuid.UUID       `json:"first_id"`
	Status      string          `json:"status"`
	Type        string          `json:"type"`
	Description *string         `json:"description,omitempty"`
	Meta        json.RawMessage `json:"meta,omitempty"`
	CreatedAt   time.Time       `json:"created_at"`
}
