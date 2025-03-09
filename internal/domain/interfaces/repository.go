package interfaces

import (
	"decard/internal/domain/entity"
	"decard/internal/domain/valueobject"
)

type ProfileRepository interface {
	FindByUUID(userUUID valueobject.UUID) (*entity.Profile, error)
	FindByTelegramID(telegramID entity.TelegramID) (*entity.Profile, error)
	Create(profile entity.Profile) error
}

type CustomerRepository interface {
	FindByProfileUUID(profileUUID valueobject.UUID) (*entity.Customer, error)
	FindByTelegramID(telegramID entity.TelegramID) (*entity.Customer, error)
	Create(customer entity.Customer) error
}

type CardRepository interface {
	GetByAccount(account valueobject.UUID) (entity.Card, error)
}

type AccountRepository interface {
	GetByCustomer(customer valueobject.UUID) (*entity.Account, error)
}
