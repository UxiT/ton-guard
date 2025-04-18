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
	Create(account entity.Account) error
}

type RefreshTokenRepository interface {
	FindByToken(token string) (*entity.RefreshToken, error)
	GetLastForProfile(profileUUID valueobject.UUID) (*entity.RefreshToken, error)
	Delete(token valueobject.RefreshToken) error
	Create(profileUUID valueobject.UUID) error
}

type TopUpRepository interface {
	Create(topUp entity.TopUp) error
	GetByUUID(uuid valueobject.UUID) (*entity.TopUp, error)
	SetStatus(uuid valueobject.UUID, status entity.TopUpStatus) error
	AddTransactionID(uuid valueobject.UUID, transactionID string) error
	GetCustomerCurrentTopUp(profileUUID valueobject.UUID) (*entity.TopUp, error)
	Close(uuid valueobject.UUID, status entity.TopUpStatus) error
}
