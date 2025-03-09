package auth

import (
	"decard/internal/domain"
	"decard/internal/domain/entity"
	"decard/internal/domain/interfaces"
	"decard/pkg/utils/hasher"
	"decard/pkg/utils/jwt"
	"log/slog"
)

type RegisterCommand struct {
	TelegramID int
	Email      string
	Password   string
}

type RegisterCommandHandler struct {
	logger            *slog.Logger
	profileRepository interfaces.ProfileRepository
}

func NewRegisterCommandHandler(
	logger *slog.Logger,
	profileRepository interfaces.ProfileRepository,
) *RegisterCommandHandler {
	return &RegisterCommandHandler{
		logger:            logger,
		profileRepository: profileRepository,
	}
}

func (h *RegisterCommandHandler) Handle(cmd RegisterCommand) (string, error) {
	const op = "application.command.register"
	logger := h.logger.With(slog.String("operation", op))

	telegramID, err := entity.NewTelegramID(cmd.TelegramID)
	if err != nil {
		return "", domain.ErrInvalidUser
	}

	email, err := entity.NewEmail(cmd.Email)
	if err != nil {
		return "", err
	}

	existingProfile, err := h.profileRepository.FindByTelegramID(telegramID)
	if err == nil && existingProfile != nil {
		return "", domain.ErrCustomerNotFound
	}

	hashedPassword, err := hasher.Hash(cmd.Password)
	if err != nil {
		logger.Error("error hashing password", slog.String("error", err.Error()))
		return "", domain.ErrInvalidCredentials
	}

	profile := entity.NewProfile(telegramID, email, hashedPassword)

	if err := h.profileRepository.Create(profile); err != nil {
		logger.Error("error creating profile", slog.String("error", err.Error()))
		return "", err
	}

	token, err := jwt.GenerateToken(profile.UUID)

	if err != nil {
		logger.Error("error generating jwt token", slog.String("error", err.Error()))
		return "", err
	}

	return token, nil
}
