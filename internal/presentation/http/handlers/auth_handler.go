package handlers

import (
	"decard/internal/application/command/auth"
	"decard/internal/presentation/http/common"
	"encoding/json"
	"net/http"
)

type LoginRequest struct {
	TelegramID int    `json:"telegram_id"`
	Password   string `json:"password"`
}

type RegisterRequest struct {
	TelegramID int    `json:"telegram_id"`
	Email      string `json:"email"`
	Password   string `json:"password"`
}

type AuthHandler struct {
	authCommandHandler     *auth.AuthenticateCommandHandler
	registerCommandHandler *auth.RegisterCommandHandler
}

func NewAuthHandler(
	authCommandHandler *auth.AuthenticateCommandHandler,
	registerCommandHandler *auth.RegisterCommandHandler,
) *AuthHandler {
	return &AuthHandler{
		authCommandHandler:     authCommandHandler,
		registerCommandHandler: registerCommandHandler,
	}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) error {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return err
	}

	token, err := h.authCommandHandler.Handle(auth.AuthenticateCommand{
		TelegramID: req.TelegramID,
		Password:   req.Password,
	})

	if err != nil {
		return err
	}

	return common.JSONResponse(w, http.StatusOK, map[string]string{"token": token})
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) error {
	var req RegisterRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return err
	}

	token, err := h.registerCommandHandler.Handle(auth.RegisterCommand{
		TelegramID: req.TelegramID,
		Email:      req.Email,
		Password:   req.Password,
	})

	if err != nil {
		return err
	}

	return common.JSONResponse(w, http.StatusOK, map[string]string{"token": token})
}
