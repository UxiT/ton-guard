package repository

import (
	"decard/internal/domain/aggregate"
	"decard/internal/domain/entity"
)

type CustomerRepository interface {
	FindByTelegramID(telegramID entity.TelegramID) (*aggregate.Customer, error)
	Create(customer aggregate.Customer) error
}
