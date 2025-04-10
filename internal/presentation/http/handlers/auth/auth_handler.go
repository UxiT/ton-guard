package auth

import (
	"context"
	"decard/internal/application/command/auth"
	"decard/internal/presentation/http/common"
	"net/http"
)

type AuthHandler struct {
	authCommandHandler     *auth.AuthenticateCommandHandler
	registerCommandHandler *auth.RegisterCommandHandler
	refreshCommandHandler  *auth.RefreshCommandHandler
}

func NewAuthHandler(
	authCommandHandler *auth.AuthenticateCommandHandler,
	registerCommandHandler *auth.RegisterCommandHandler,
	refreshCommandHandler *auth.RefreshCommandHandler,
) *AuthHandler {
	return &AuthHandler{
		authCommandHandler:     authCommandHandler,
		registerCommandHandler: registerCommandHandler,
		refreshCommandHandler:  refreshCommandHandler,
	}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r any) error {
	req := r.(*LoginRequest)

	response, err := h.authCommandHandler.Handle(context.Background(), auth.AuthenticateCommand{
		TelegramID: req.TelegramID,
		Password:   req.Password,
	})

	if err != nil {
		return err
	}

	return common.JSONResponse(w, http.StatusOK, response)
}

func (h *AuthHandler) Register(w http.ResponseWriter, r any) error {
	req := r.(*RegisterRequest)

	response, err := h.registerCommandHandler.Handle(context.Background(), auth.RegisterCommand{
		TelegramID: req.TelegramID,
		Email:      req.Email,
		Password:   req.Password,
	})

	if err != nil {
		return err
	}

	return common.JSONResponse(w, http.StatusOK, response)
}

func (h *AuthHandler) Refresh(w http.ResponseWriter, r any) error {
	req := r.(*RefreshRequest)

	response, err := h.refreshCommandHandler.Handle(context.Background(), auth.RefreshCommand{
		RefreshToken: req.RefreshToken,
	})

	if err != nil {
		return err
	}

	return common.JSONResponse(w, http.StatusOK, response)
}
