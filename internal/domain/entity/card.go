package entity

import (
	"github.com/google/uuid"
)

type Card struct {
	UUID         uuid.UUID
	ExternalUUID uuid.UUID
	AccountUUID  uuid.UUID
}
