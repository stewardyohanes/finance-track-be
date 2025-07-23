package users

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type (
	SignInRequest struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	SignInResponse struct {
		Data struct {
			AccessToken  string `json:"access_token"`
			RefreshToken string `json:"refresh_token"`
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
			ID           uuid.UUID `json:"id"`
			Username     string    `json:"username"`
			Email        string    `json:"email"`
			CreatedAt    time.Time `json:"created_at"`
			UpdatedAt    time.Time `json:"updated_at"`
			AccessToken  string    `json:"access_token"`
			RefreshToken string    `json:"refresh_token"`
		} `json:"data"`
		Message string `json:"message"`
	}

	RefreshTokenRequest struct {
		RefreshToken string `json:"refresh_token" binding:"required"`
	}

	RefreshTokenResponse struct {
		Data struct {
			AccessToken  string `json:"access_token"`
			RefreshToken string `json:"refresh_token"`
		} `json:"data"`
		Message string `json:"message"`
	}
)

type (
	User struct {
		ID           uuid.UUID  `gorm:"type:uuid;primaryKey"`
		Username     string     `gorm:"unique;not null"`
		Email        string     `gorm:"unique;not null"`
		Password     string     `gorm:"not null"`
		RefreshToken *string    `gorm:"type:text"`
		TokenExpiry  *time.Time `gorm:"type:timestamp"`
		CreatedAt    time.Time  `gorm:"autoCreateTime"`
		UpdatedAt    time.Time  `gorm:"autoUpdateTime"`
	}
)

// BeforeCreate hook untuk generate UUID otomatis
func (u *User) BeforeCreate(tx *gorm.DB) error {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	return nil
}