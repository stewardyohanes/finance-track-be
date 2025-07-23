package users

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"github.com/stewardyohanes/finance-tracker/internal/models/users"
	"gorm.io/gorm"
)

func (r *repository) CreateUser(user *users.User) error {
	if err := r.db.Create(user).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			log.Error().Err(err).Msg("Email or username already exists")
			return errors.New("email or username already exists")
		}
		return err
	}
	return nil
}

func (r *repository) GetUserByEmail(email string) (*users.User, error) {
	var user users.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Error().Err(err).Msg("User Email not found")
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *repository) GetUserByUsername(username string) (*users.User, error) {
	var user users.User
	if err := r.db.Where("username = ?", username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Error().Err(err).Msg("User Username not found")
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *repository) GetUserByID(id uuid.UUID) (*users.User, error) {
	var user users.User
	if err := r.db.Where("id = ?", id).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Error().Err(err).Msg("User ID not found")
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *repository) UpdateRefreshToken(userID uuid.UUID, refreshToken string, expiry time.Time) error {
	if err := r.db.Model(&users.User{}).Where("id = ?", userID).Updates(map[string]interface{}{
		"refresh_token": refreshToken,
		"token_expiry":  expiry,
	}).Error; err != nil {
		log.Error().Err(err).Msg("Failed to update refresh token")
		return err
	}
	return nil
}

func (r *repository) GetUserByRefreshToken(refreshToken string) (*users.User, error) {
	var user users.User
	if err := r.db.Where("refresh_token = ? AND token_expiry > ?", refreshToken, time.Now()).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Error().Err(err).Msg("Invalid or expired refresh token")
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *repository) ClearRefreshToken(userID uuid.UUID) error {
	if err := r.db.Model(&users.User{}).Where("id = ?", userID).Updates(map[string]interface{}{
		"refresh_token": nil,
		"token_expiry":  nil,
	}).Error; err != nil {
		log.Error().Err(err).Msg("Failed to clear refresh token")
		return err
	}
	return nil
}