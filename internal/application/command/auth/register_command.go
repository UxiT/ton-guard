package auth

import (
	"context"
	"decard/internal/domain"
	"decard/internal/domain/entity"
	"decard/internal/domain/interfaces"
	"decard/pkg/utils/hasher"
	"log/slog"
)

type RegisterCommand struct {
	TelegramID int
	Email      string
	Password   string
}

type RegisterCommandHandler struct {
	logger             *slog.Logger
	profileRepository  interfaces.ProfileRepository
	generateJWTCommand *GenerateJWTCommandHandler
}

func NewRegisterCommandHandler(
	logger *slog.Logger,
	profileRepository interfaces.ProfileRepository,
	generateJWTCommand *GenerateJWTCommandHandler,
) *RegisterCommandHandler {
	return &RegisterCommandHandler{
		logger:             logger,
		profileRepository:  profileRepository,
		generateJWTCommand: generateJWTCommand,
	}
}

func (h *RegisterCommandHandler) Handle(ctx context.Context, cmd RegisterCommand) (GenerateJWTResponse, error) {
	const op = "application.command.register"
	logger := h.logger.With(slog.String("operation", op))

	telegramID, err := entity.NewTelegramID(cmd.TelegramID)
	if err != nil {
		return GenerateJWTResponse{}, domain.ErrInvalidUser
	}

	email, err := entity.NewEmail(cmd.Email)
	if err != nil {
		return GenerateJWTResponse{}, err
	}

	existingProfile, err := h.profileRepository.FindByTelegramID(telegramID)
	if err == nil && existingProfile != nil {
		return GenerateJWTResponse{}, domain.ErrCustomerNotFound
	}

	hashedPassword, err := hasher.Hash(cmd.Password)
	if err != nil {
		logger.Error("error hashing password", slog.String("error", err.Error()))
		return GenerateJWTResponse{}, domain.ErrInvalidCredentials
	}

	profile := entity.NewProfile(telegramID, email, hashedPassword)

	if err := h.profileRepository.Create(profile); err != nil {
		logger.Error("error creating profile", slog.String("error", err.Error()))
		return GenerateJWTResponse{}, err
	}

	return h.generateJWTCommand.Handle(ctx, GenerateJWTCommand{profile.UUID})
}
