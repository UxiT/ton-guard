package entity

import (
	"decard/internal/domain/valueobject"
	"encoding/json"
	"time"
)

// BatchFinal represents the final state of a batch.
type BatchFinal struct {
	UUID        valueobject.UUID `json:"uuid"`
	Status      string           `json:"status"`
	Type        string           `json:"type"`
	Description *string          `json:"description,omitempty"`
	Meta        json.RawMessage  `json:"meta,omitempty"`
	CreatedAt   time.Time        `json:"created_at"`
	UpdatedAt   time.Time        `json:"updated_at"`
}
