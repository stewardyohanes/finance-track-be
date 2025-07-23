package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Payload struct {
	ID        uuid.UUID `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	SecretKey string `json:"secret_key"`
}

func CreateToken(payload *Payload) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":         payload.ID.String(),
		"username":   payload.Username,
		"email":      payload.Email,
		"exp":        time.Now().Add(time.Hour * 24).Unix(),
	})
	
	key := []byte(payload.SecretKey)
	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ValidateToken(tokenString, secretKey string) (uuid.UUID, string, string, error) {
	key := []byte(secretKey)
	claims := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		return uuid.Nil, "", "", err
	}

	if !token.Valid {
		return uuid.Nil, "", "", err
	}

	return uuid.MustParse(claims["id"].(string)), claims["username"].(string), claims["email"].(string), nil
}