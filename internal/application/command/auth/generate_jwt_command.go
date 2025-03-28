package auth

import (
	"context"
	"decard/internal/domain/service"
	"decard/internal/domain/valueobject"
	"decard/pkg/utils/jwt"
	"github.com/rs/zerolog"
)

type GenerateJWTCommand struct {
	ProfileUUID valueobject.UUID
}

type GenerateJWTResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

type GenerateJWTCommandHandler struct {
	logger      *zerolog.Logger
	authService *service.AuthService
}

func NewGenerateJWTCommandHandler(
	logger *zerolog.Logger,
	authService *service.AuthService,
) *GenerateJWTCommandHandler {
	return &GenerateJWTCommandHandler{
		logger:      logger,
		authService: authService,
	}
}

func (h GenerateJWTCommandHandler) Handle(ctx context.Context, cmd GenerateJWTCommand) (GenerateJWTResponse, error) {
	token, err := jwt.GenerateToken(cmd.ProfileUUID)

	if err != nil {
		h.logger.Error().Err(err).Msg("error generating jwt token")

		return GenerateJWTResponse{}, err
	}

	refreshToken, err := h.authService.GenerateRefreshToken(cmd.ProfileUUID)

	if err != nil {
		h.logger.Error().Err(err).Msg("error generating refresh token")

		return GenerateJWTResponse{}, err
	}

	return GenerateJWTResponse{
		Token:        token,
		RefreshToken: refreshToken.UUID.String(),
	}, nil
}
