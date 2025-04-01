package service

import (
	"decard/internal/domain/entity"
	"decard/internal/domain/interfaces"
	"decard/internal/domain/valueobject"
	"github.com/rs/zerolog"
)

type AuthService struct {
	logger           *zerolog.Logger
	refreshTokenRepo interfaces.RefreshTokenRepository
}

func NewAuthService(logger *zerolog.Logger, refreshTokenRepo interfaces.RefreshTokenRepository) *AuthService {
	return &AuthService{
		logger:           logger,
		refreshTokenRepo: refreshTokenRepo,
	}
}

func (s AuthService) GenerateRefreshToken(profile valueobject.UUID) (*entity.RefreshToken, error) {
	existing, err := s.refreshTokenRepo.GetLastForProfile(profile)

	if err != nil && existing == nil {
		s.logger.Error().Str("profile", profile.String()).Err(err).Msg("Error getting last refresh token for profile")
	}

	if existing != nil {
		err = s.refreshTokenRepo.Delete(existing.UUID)

		if err != nil {
			s.logger.Error().Err(err).Msg("Error deleting refresh token")
			return nil, err
		}
	}

	if err := s.refreshTokenRepo.Create(profile); err != nil {
		s.logger.Error().Str("profile", profile.String()).Err(err).Msg("Error creating refresh token")

		return nil, err
	}

	return s.refreshTokenRepo.GetLastForProfile(profile)
}
