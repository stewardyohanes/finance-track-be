package users

import (
	"errors"

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