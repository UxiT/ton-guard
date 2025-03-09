package entity

import (
	"decard/internal/domain/valueobject"
	"time"
)

type Customer struct {
	UUID        valueobject.UUID
	ProfileUUID valueobject.UUID

	CreatedAt time.Time
	UpdatedAt time.Time
}
