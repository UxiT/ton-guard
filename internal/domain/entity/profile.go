package entity

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrInvalidUUID        = errors.New("invalid UUID provided")
	ErrInvalidPassword    = errors.New("invalid password")
	ErrInvalidCredentials = errors.New("invalid Credentials")
)

type Profile struct {
	UUID         uuid.UUID
	TelegramID   TelegramID
	Email        Email
	PasswordHash string
}

func NewProfile(telegramID TelegramID, email Email, password string) Profile {
	return Profile{
		UUID:         uuid.New(),
		TelegramID:   telegramID,
		Email:        email,
		PasswordHash: password,
	}
}

type TelegramID int

func NewTelegramID(i int) (TelegramID, error) {
	return TelegramID(i), nil
}

func (i TelegramID) Int() int {
	return int(i)
}

type Email string

func NewEmail(i string) (Email, error) {
	return Email(i), nil
}

func (i Email) String() string {
	return string(i)
}
