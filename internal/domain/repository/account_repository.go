package repository

import (
	"decard/internal/domain/entity"

	"github.com/google/uuid"
)

type AccountRepository interface {
	GetByCustomer(customer uuid.UUID) (entity.Account, error)
}
