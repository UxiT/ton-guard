package entity

import (
	"decard/internal/domain/valueobject"
	"encoding/json"
	"time"
)

// Batch represents a batch record.
type Batch struct {
	UUID        valueobject.UUID `json:"uuid"`
	FirstID     valueobject.UUID `json:"first_id"`
	Status      string           `json:"status"`
	Type        string           `json:"type"`
	Description *string          `json:"description,omitempty"`
	Meta        json.RawMessage  `json:"meta,omitempty"`
	CreatedAt   time.Time        `json:"created_at"`
}
