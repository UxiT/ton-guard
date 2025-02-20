package handlers

import (
	"decard/internal/domain/entity"
	"decard/internal/domain/service"
	"encoding/json"
	"errors"
	"net/http"
)

// LoginRequest and RegisterRequest are presentation-layer DTOs,
// providing a contract for parsing incoming JSON.
type LoginRequest struct {
	TelegramID int    `json:"telegram_id"`
	Password   string `json:"password"`
}

type RegisterRequest struct {
	TelegramID int    `json:"telegram_id"`
	Email      string `json:"email"`
	Password   string `json:"password"`
}

// AuthHandler is part of the adapter layer. It only translates HTTP messages
// into application commands and delegates the business logic to the domain service.
type AuthHandler struct {
	authService service.AuthService // Ideally an interface from the domain/application layer.
}

// NewAuthHandler constructs an AuthHandler with the provided domain service.
func NewAuthHandler(authService service.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

// Login parses a login command from the HTTP request and delegates to the domain service.
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		jsonError(w, "Invalid request", http.StatusBadRequest)

		return
	}

	telegramID, err := entity.NewTelegramID(req.TelegramID)
	if err != nil {
		jsonError(w, err.Error(), http.StatusUnprocessableEntity)

		return
	}

	token, err := h.authService.Authenticate(telegramID, req.Password)
	if err != nil {
		if errors.Is(err, entity.ErrInvalidCredentials) {
			jsonError(w, err.Error(), http.StatusUnauthorized)
		} else {
			jsonError(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	jsonResponse(w, map[string]string{"token": token}, http.StatusOK)
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		jsonError(w, "Invalid request", http.StatusBadRequest)

		return
	}

	telegramID, err := entity.NewTelegramID(req.TelegramID)
	if err != nil {
		jsonError(w, err.Error(), http.StatusUnprocessableEntity)

		return
	}

	email, err := entity.NewEmail(req.Email)
	if err != nil {
		jsonError(w, err.Error(), http.StatusUnprocessableEntity)

		return
	}

	token, err := h.authService.Register(telegramID, email, req.Password)
	if err != nil {
		jsonError(w, err.Error(), http.StatusInternalServerError)

		return
	}

	jsonResponse(w, map[string]string{"token": token}, http.StatusOK)
}

// jsonResponse is a helper to send JSON responses with a given status.
func jsonResponse(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	// Ignoring encoding errors for simplicity.
	json.NewEncoder(w).Encode(data)
}

// jsonError sends an error message as JSON with the supplied HTTP status code.
func jsonError(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}
