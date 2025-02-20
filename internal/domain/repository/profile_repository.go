package repository

import (
	"decard/internal/domain/entity"
)

type ProfileRepository interface {
	FindByTelegramID(telegramID entity.TelegramID) (*entity.Profile, error)
	Create(profile entity.Profile) error
}
