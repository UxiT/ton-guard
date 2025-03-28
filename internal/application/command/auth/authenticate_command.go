package auth

import (
	"context"
	"decard/internal/domain"
	"decard/internal/domain/entity"
	"decard/internal/domain/interfaces"
	"decard/pkg/utils/hasher"
	"github.com/rs/zerolog"
)

type AuthenticateCommand struct {
	TelegramID int
	Password   string
}

type AuthenticateCommandHandler struct {
	logger             *zerolog.Logger
	profileRepository  interfaces.ProfileRepository
	generateJWTCommand *GenerateJWTCommandHandler
}

func NewAuthenticateCommandHandler(
	logger *zerolog.Logger,
	profileRepository interfaces.ProfileRepository,
	generateJWTCommand *GenerateJWTCommandHandler,
) *AuthenticateCommandHandler {
	return &AuthenticateCommandHandler{
		logger:             logger,
		profileRepository:  profileRepository,
		generateJWTCommand: generateJWTCommand,
	}
}

func (h AuthenticateCommandHandler) Handle(ctx context.Context, cmd AuthenticateCommand) (GenerateJWTResponse, error) {
	const op = "application.command.authenticate"

	logger := h.logger.With().Str("op", op).Logger()

	telegramID, err := entity.NewTelegramID(cmd.TelegramID)

	if err != nil {
		return GenerateJWTResponse{}, err
	}

	profile, err := h.profileRepository.FindByTelegramID(telegramID)
	if err != nil {
		logger.Error().Int("telegram_id", telegramID.Int()).Err(err).Msg("error finding customer")

		return GenerateJWTResponse{}, domain.ErrCustomerNotFound
	}

	if !hasher.VerifyPassword(profile.PasswordHash, cmd.Password) {
		logger.Error().Err(err).Msg("error when verifying password")

		return GenerateJWTResponse{}, domain.ErrInvalidCredentials
	}

	return h.generateJWTCommand.Handle(ctx, GenerateJWTCommand{ProfileUUID: profile.UUID})
}
