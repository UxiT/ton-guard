package entity

import "decard/internal/domain/valueobject"

type RefreshToken struct {
	UUID        valueobject.RefreshToken
	ProfileUUID valueobject.UUID
}
