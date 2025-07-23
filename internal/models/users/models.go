package users

import (
	"time"

	"github.com/google/uuid"
)

type (
	UserModel struct {
		ID        uuid.UUID `gorm:"primaryKey"`
		Username  string    `gorm:"unique;not null"`
		Email     string    `gorm:"unique;not null"`
		Password  string    `gorm:"not null"`
		CreatedAt time.Time `gorm:"autoCreateTime"`
		UpdatedAt time.Time `gorm:"autoUpdateTime"`
	}
)