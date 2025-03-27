package service

import (
	"decard/internal/domain/entity"
	"decard/internal/domain/interfaces"
	"decard/internal/domain/valueobject"
	"log/slog"
)

type AuthService struct {
	logger           *slog.Logger
	refreshTokenRepo interfaces.RefreshTokenRepository
}

func NewAuthService(logger *slog.Logger, refreshTokenRepo interfaces.RefreshTokenRepository) *AuthService {
	return &AuthService{
		logger:           logger,
		refreshTokenRepo: refreshTokenRepo,
	}
}

func (s AuthService) GenerateRefreshToken(profile valueobject.UUID) (*entity.RefreshToken, error) {
	existing, err := s.refreshTokenRepo.GetLastForProfile(profile)
	if err != nil {
		s.logger.Error("Error getting last refresh token for profile", slog.String("profile", profile.String()), slog.String("error", err.Error()))
		return nil, err
	}

	if existing != nil {
		err = s.refreshTokenRepo.Delete(existing.UUID)

		if err != nil {
			s.logger.Error("Error deleting refresh token", slog.String("error", err.Error()))
			return nil, err
		}
	}

	if err := s.refreshTokenRepo.Create(profile); err != nil {
		s.logger.Error("Error creating refresh token", slog.String("profile", profile.String()), slog.String("error", err.Error()))

		return nil, err
	}

	return s.refreshTokenRepo.GetLastForProfile(profile)
}
