package entity

import (
	"decard/internal/domain/valueobject"
)

type Card struct {
	UUID         valueobject.UUID
	ExternalUUID valueobject.UUID
	AccountUUID  valueobject.UUID
}
