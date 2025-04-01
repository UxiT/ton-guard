package auth

import (
	"context"
	"decard/internal/domain/interfaces"
	"github.com/rs/zerolog"
)

type RefreshCommand struct {
	RefreshToken string
}

type RefreshCommandHandler struct {
	logger                 *zerolog.Logger
	refreshTokenRepository interfaces.RefreshTokenRepository
	generateJWTCommand     *GenerateJWTCommandHandler
}

func NewRefreshCommandHandler(
	logger *zerolog.Logger,
	refreshTokenRepository interfaces.RefreshTokenRepository,
	generateJWTCommand *GenerateJWTCommandHandler,
) *RefreshCommandHandler {
	return &RefreshCommandHandler{
		logger:                 logger,
		refreshTokenRepository: refreshTokenRepository,
		generateJWTCommand:     generateJWTCommand,
	}
}

func (h RefreshCommandHandler) Handle(ctx context.Context, cmd RefreshCommand) (GenerateJWTResponse, error) {
	token, err := h.refreshTokenRepository.FindByToken(cmd.RefreshToken)

	if err != nil {
		h.logger.Error().Err(err).Str("token", cmd.RefreshToken).Msg("refresh token not found")
		return GenerateJWTResponse{}, err
	}

	if err = h.refreshTokenRepository.Delete(token.UUID); err != nil {
		return GenerateJWTResponse{}, err
	}

	return h.generateJWTCommand.Handle(ctx, GenerateJWTCommand{token.ProfileUUID})
}
