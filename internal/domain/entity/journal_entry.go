package entity

import (
	"decard/internal/domain/valueobject"
	"encoding/json"
	"time"

	"github.com/shopspring/decimal"
)

// JournalEntry represents an entry in the journal.
type JournalEntry struct {
	UUID        valueobject.UUID  `json:"uuid"`
	Type        string            `json:"type"`
	Amount      decimal.Decimal   `json:"amount"`
	Description *string           `json:"description,omitempty"`
	Meta        json.RawMessage   `json:"meta,omitempty"`
	AccountID   *valueobject.UUID `json:"account_id,omitempty"`
	BatchID     *valueobject.UUID `json:"batch_id,omitempty"`
	CreatedAt   time.Time         `json:"created_at"`
}
