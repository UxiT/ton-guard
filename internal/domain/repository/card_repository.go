package repository

import (
	"decard/internal/domain/entity"

	"github.com/google/uuid"
)

type CardRepository interface {
	GetByCustomer(customer uuid.UUID) ([]entity.Card, error)
}
