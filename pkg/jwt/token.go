package jwt

import (
	"crypto/rand"
	"encoding/hex"
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

type TokenPair struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// CreateAccessToken membuat JWT access token dengan durasi pendek (15 menit)
func CreateAccessToken(payload *Payload) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":         payload.ID.String(),
		"username":   payload.Username,
		"email":      payload.Email,
		"exp":        time.Now().Add(time.Minute * 15).Unix(), // 15 menit
		"type":       "access",
	})
	
	key := []byte(payload.SecretKey)
	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// CreateRefreshToken membuat refresh token random string
func CreateRefreshToken() (string, error) {
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

// CreateTokenPair membuat pasangan access token dan refresh token
func CreateTokenPair(payload *Payload) (*TokenPair, error) {
	accessToken, err := CreateAccessToken(payload)
	if err != nil {
		return nil, err
	}

	refreshToken, err := CreateRefreshToken()
	if err != nil {
		return nil, err
	}

	return &TokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

// ValidateAccessToken validasi JWT access token
func ValidateAccessToken(tokenString, secretKey string) (uuid.UUID, string, string, error) {
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

	// Validasi type token
	if tokenType, ok := claims["type"].(string); !ok || tokenType != "access" {
		return uuid.Nil, "", "", jwt.ErrTokenInvalidClaims
	}

	return uuid.MustParse(claims["id"].(string)), claims["username"].(string), claims["email"].(string), nil
}

// Legacy function untuk backward compatibility
func CreateToken(payload *Payload) (string, error) {
	return CreateAccessToken(payload)
}

// Legacy function untuk backward compatibility
func ValidateToken(tokenString, secretKey string) (uuid.UUID, string, string, error) {
	return ValidateAccessToken(tokenString, secretKey)
}