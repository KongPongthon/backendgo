package utils

import (
	"time"

	"backendgo/config"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(userID string) (string, error) {
    claims := jwt.MapClaims{
        "user_id": userID,
        "exp":     time.Now().Add(time.Hour * 72).Unix(),
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(config.JwtSecret)
}
