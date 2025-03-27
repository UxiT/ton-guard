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

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) error {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return err
	}

	response, err := h.authCommandHandler.Handle(r.Context(), auth.AuthenticateCommand{
		TelegramID: req.TelegramID,
		Password:   req.Password,
	})

	if err != nil {
		return err
	}

	return common.JSONResponse(w, http.StatusOK, response)
}

type RegisterRequest struct {
	TelegramID int    `json:"telegram_id"`
	Email      string `json:"email"`
	Password   string `json:"password"`
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) error {
	var req RegisterRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return err
	}

	response, err := h.registerCommandHandler.Handle(r.Context(), auth.RegisterCommand{
		TelegramID: req.TelegramID,
		Email:      req.Email,
		Password:   req.Password,
	})

	if err != nil {
		return err
	}

	return common.JSONResponse(w, http.StatusOK, response)
}

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token"`
}

func (h *AuthHandler) Refresh(w http.ResponseWriter, r *http.Request) error {
	var req RefreshRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return err
	}

	response, err := h.refreshCommandHandler.Handle(r.Context(), auth.RefreshCommand{RefreshToken: req.RefreshToken})

	if err != nil {
		return err
	}

	return common.JSONResponse(w, http.StatusOK, response)
}
