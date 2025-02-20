package entity

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// JournalEntry represents an entry in the journal.
type JournalEntry struct {
	UUID        uuid.UUID       `json:"uuid"`
	Type        string          `json:"type"`
	Amount      decimal.Decimal `json:"amount"`
	Description *string         `json:"description,omitempty"`
	Meta        json.RawMessage `json:"meta,omitempty"`
	// Both AccountID and BatchID are represented as pointers as these fields are not marked as NOT NULL in the migration.
	AccountID *uuid.UUID `json:"account_id,omitempty"`
	BatchID   *uuid.UUID `json:"batch_id,omitempty"`
	CreatedAt time.Time  `json:"created_at"`
}
