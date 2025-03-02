package interfaces

import (
	"decard/internal/domain/aggregate"
	"decard/internal/domain/entity"
	"github.com/google/uuid"
)

type ProfileRepository interface {
	FindByTelegramID(telegramID entity.TelegramID) (*entity.Profile, error)
	Create(profile entity.Profile) error
}

type CustomerRepository interface {
	FindByTelegramID(telegramID entity.TelegramID) (*aggregate.Customer, error)
	Create(customer aggregate.Customer) error
}

type CardRepository interface {
	GetByAccount(account uuid.UUID) (entity.Card, error)
}

type AccountRepository interface {
	GetByCustomer(customer uuid.UUID) (aggregate.Account, error)
}
