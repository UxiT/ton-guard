package entity

import (
	"decard/internal/domain/valueobject"
	"encoding/json"
	"time"

	"github.com/shopspring/decimal"
)

// ExchangeRate represents an exchange rate entry.
type ExchangeRate struct {
	UUID           valueobject.UUID `json:"uuid"`
	IsoCode        string           `json:"iso_code"`
	Rate           decimal.Decimal  `json:"rate"`
	RelatedIsoCode string           `json:"related_iso_code"`
	Meta           json.RawMessage  `json:"meta,omitempty"`
	CreatedAt      time.Time        `json:"created_at"`
}
