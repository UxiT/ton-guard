package auth

import (
	"context"
	"decard/internal/domain"
	"decard/internal/domain/entity"
	"decard/internal/domain/interfaces"
	"decard/pkg/utils/hasher"
	"log/slog"
)

type AuthenticateCommand struct {
	TelegramID int
	Password   string
}

type AuthenticateCommandHandler struct {
	logger             *slog.Logger
	profileRepository  interfaces.ProfileRepository
	generateJWTCommand *GenerateJWTCommandHandler
}

func NewAuthenticateCommandHandler(
	logger *slog.Logger,
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

	logger := h.logger.With(slog.String("op", op))

	telegramID, err := entity.NewTelegramID(cmd.TelegramID)

	if err != nil {
		return GenerateJWTResponse{}, err
	}

	profile, err := h.profileRepository.FindByTelegramID(telegramID)
	if err != nil {
		logger.Error("error finding customer", slog.Int("telegramID", cmd.TelegramID), slog.String("error", err.Error()))

		return GenerateJWTResponse{}, domain.ErrCustomerNotFound
	}

	if !hasher.VerifyPassword(profile.PasswordHash, cmd.Password) {
		logger.Error("error when verifying password")

		return GenerateJWTResponse{}, domain.ErrInvalidCredentials
	}

	return h.generateJWTCommand.Handle(ctx, GenerateJWTCommand{ProfileUUID: profile.UUID})
}
