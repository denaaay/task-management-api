package utils

import (
	"fmt"
	"time"

	"github.com/denaaay/task-management-api/model"
	"github.com/golang-jwt/jwt"
)

func GenerateToken(user model.UserPayload) (string, error) {
	claims := jwt.MapClaims{
		"user_id":  user.Id,
		"username": user.Name,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("rahasia"))
	if err != nil {
		return "", fmt.Errorf("error generating token: %v", err)
	}

	return tokenString, nil
}

func ValidateToken(token string) (any, error) {
	tkn, err := jwt.Parse(token, func(jwtToken *jwt.Token) (any, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected method: %s", jwtToken.Header["alg"])
		}

		return []byte("rahasia"), nil
	})

	if err != nil {
		return nil, fmt.Errorf("invalidate token: %w", err)
	}

	claims, ok := tkn.Claims.(jwt.MapClaims)
	if !ok || !tkn.Valid {
		return nil, fmt.Errorf("invalid token claim")
	}

	return claims["user_id"], nil
}
