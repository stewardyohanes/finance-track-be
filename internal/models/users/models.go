package users

import (
	"time"

	"github.com/google/uuid"
)

type (
	SignInRequest struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	SignInResponse struct {
		Data struct {
			Token string `json:"token"`
		} `json:"data"`
		Message string `json:"message"`
	}

	SignUpRequest struct {
		Username string `json:"username" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	SignUpResponse struct {
		Data struct {
			ID        uuid.UUID `json:"id"`
			Username  string    `json:"username"`
			Email     string    `json:"email"`
			CreatedAt time.Time `json:"created_at"`
			UpdatedAt time.Time `json:"updated_at"`
			Token     string    `json:"token"`
		} `json:"data"`
		Message string `json:"message"`
	}
)

type (
	User struct {
		ID        uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
		Username  string    `gorm:"unique;not null"`
		Email     string    `gorm:"unique;not null"`
		Password  string    `gorm:"not null"`
		CreatedAt time.Time `gorm:"autoCreateTime"`
		UpdatedAt time.Time `gorm:"autoUpdateTime"`
		
	}
)