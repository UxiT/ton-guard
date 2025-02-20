package middleware

import (
	"context"
	"decard/pkg/utils/jwt"
	"net/http"
	"strings"
)

type TelegramIDContextKey string

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "No authorization header", http.StatusUnauthorized)
			return
		}

		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
		telegramID, err := jwt.ValidateToken(tokenString)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), TelegramIDContextKey(telegramID), telegramID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
