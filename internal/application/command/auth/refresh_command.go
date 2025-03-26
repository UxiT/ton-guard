package auth

import (
	"context"
	"decard/internal/domain/interfaces"
	"decard/internal/domain/service"
	"log/slog"
)

type RefreshCommand struct {
	RefreshToken string
}

type RefreshCommandHandler struct {
	logger                 *slog.Logger
	refreshTokenRepository interfaces.RefreshTokenRepository
	authService            service.AuthService
}

func NewRefreshCommandHandler(logger *slog.Logger) RefreshCommandHandler {
	return RefreshCommandHandler{
		logger: logger,
	}
}

func (h RefreshCommandHandler) Handle(ctx context.Context, cmd RefreshCommand) (string, error) {
	token, err := h.refreshTokenRepository.FindByToken(cmd.RefreshToken)

	if err != nil {
		return "", err
	}

	err = h.refreshTokenRepository.Delete(token.UUID)

	if err != nil {
		return "", err
	}

	newToken, err :=  h.authService.GenerateRefreshToken(cmd.)
}
