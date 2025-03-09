package auth

import (
	"decard/internal/domain"
	"decard/internal/domain/entity"
	"decard/internal/domain/interfaces"
	"decard/pkg/utils/hasher"
	"decard/pkg/utils/jwt"
	"log/slog"
)

type AuthenticateCommand struct {
	TelegramID int
	Password   string
}

type AuthenticateCommandHandler struct {
	logger            *slog.Logger
	profileRepository interfaces.ProfileRepository
}

func NewAuthenticateCommandHandler(
	logger *slog.Logger,
	profileRepository interfaces.ProfileRepository,
) *AuthenticateCommandHandler {
	return &AuthenticateCommandHandler{
		logger:            logger,
		profileRepository: profileRepository,
	}
}

func (h AuthenticateCommandHandler) Handle(cmd AuthenticateCommand) (string, error) {
	const op = "application.command.authenticate"

	logger := h.logger.With(slog.String("op", op))

	telegramID, err := entity.NewTelegramID(cmd.TelegramID)

	if err != nil {
		return "", err
	}

	profile, err := h.profileRepository.FindByTelegramID(telegramID)
	if err != nil {
		logger.Error("error finding customer", slog.Int("telegramID", cmd.TelegramID), slog.String("error", err.Error()))

		return "", domain.ErrCustomerNotFound
	}

	if !hasher.VerifyPassword(profile.PasswordHash, cmd.Password) {
		logger.Error("error when verifying password")

		return "", domain.ErrInvalidCredentials
	}

	token, err := jwt.GenerateToken(profile.UUID)
	if err != nil {
		logger.Error("error generating jwt token", slog.String("error", err.Error()))
		return "", err
	}

	return token, nil
}
