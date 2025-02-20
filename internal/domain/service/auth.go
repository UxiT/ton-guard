package service

import (
	"decard/internal/domain/entity"
	"decard/internal/domain/repository"
	"decard/pkg/utils/hasher"
	"decard/pkg/utils/jwt"
	"errors"
)

var ErrCustomerNotFound = errors.New("customer not found")
var ErrCustomerAlreadyExists = errors.New("customer already exists")
var ErrInvalidPassword = errors.New("invalid password")

// AuthService provides authentication logic for customers.
type AuthService struct {
	repo repository.ProfileRepository
}

// NewAuthService creates a new instance of AuthService.
func NewAuthService(repo repository.ProfileRepository) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (s *AuthService) Authenticate(telegramID entity.TelegramID, password string) (string, error) {
	profile, err := s.repo.FindByTelegramID(telegramID)
	if err != nil {
		return "", ErrCustomerNotFound
	}

	if !hasher.VerifyPassword(profile.PasswordHash, password) {
		return "", ErrInvalidPassword
	}

	token, err := jwt.GenerateToken(profile.UUID)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *AuthService) Register(telegramID entity.TelegramID, email entity.Email, password string) (string, error) {
	// Check if customer already exists
	existingProfile, err := s.repo.FindByTelegramID(telegramID)
	if err == nil && existingProfile != nil {
		return "", ErrCustomerAlreadyExists
	}

	// Hash the password
	hashedPassword, err := hasher.Hash(password)
	if err != nil {
		return "", ErrInvalidPassword
	}

	profile := entity.NewProfile(telegramID, email, hashedPassword)
	if err := s.repo.Create(profile); err != nil {
		return "", err
	}

	token, err := jwt.GenerateToken(profile.UUID)
	if err != nil {
		return "", err
	}

	return token, nil
}
