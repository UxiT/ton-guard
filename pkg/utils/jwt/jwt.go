package jwt

import (
	"decard/config"
	"decard/internal/domain"
	"decard/internal/domain/valueobject"
	"fmt"
	"github.com/golang-jwt/jwt"
	"log"
	"time"
)

type Claims struct {
	ProfileUUID valueobject.UUID `json:"profile_uuid"`
	jwt.StandardClaims
}

func GenerateToken(profileUUID valueobject.UUID) (string, error) {
	claims := &Claims{
		ProfileUUID: profileUUID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(config.Cfg.JWTSecret))
}

func ValidateToken(tokenString string) (valueobject.UUID, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(config.Cfg.JWTSecret), nil
	})

	if err != nil {
		log.Printf("error: %s", err.Error())
		return valueobject.UUID{}, domain.ErrInvalidJWT
	}

	if !token.Valid {
		return valueobject.UUID{}, domain.ErrInvalidJWT
	}

	return claims.ProfileUUID, nil
}
