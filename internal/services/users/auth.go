package users

import (
	"context"
	"errors"

	"github.com/rs/zerolog/log"
	"github.com/stewardyohanes/finance-tracker/internal/models/users"
	"github.com/stewardyohanes/finance-tracker/pkg/jwt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func (s *service) SignIn(ctx context.Context, req *users.SignInRequest) (string, error) {
	// Check if the user exists
	user, err := s.usersRepository.GetUserByUsername(req.Username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", errors.New("user not found")
		}
		log.Error().Err(err).Msg("Failed to get user")
		return "", err
	}
	
	// Check if the password is correct
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		log.Error().Err(err).Msg("Failed to compare password")
		return "", errors.New("invalid password")
	}

	token, err := jwt.CreateToken(&jwt.Payload{
		ID: user.ID,
		Username: user.Username,
		Email: user.Email,
		SecretKey: s.config.JWTKey,
	})
	if err != nil {
		log.Error().Err(err).Msg("Failed to create token")
		return "", err
	}

	return token, nil
}

func (s *service) SignUp(ctx context.Context, req *users.SignUpRequest) (string, error) {
	// Check if the user already exists
	existingEmail, err := s.usersRepository.GetUserByEmail(req.Email)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get user by email")
		return "", err
	}

	existingUsername, err := s.usersRepository.GetUserByUsername(req.Username)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get user by username")
		return "", err
	}

	if existingEmail != nil {
		return "", errors.New("email already exists")
	}

	if existingUsername != nil {	
		return "", errors.New("username already exists")
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Error().Err(err).Msg("Failed to hash password")
		return "", err
	}

	// Create the user
	user := &users.User{
		Username: req.Username,
		Email: req.Email,
		Password: string(hashedPassword),
	}

	err = s.usersRepository.CreateUser(user)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create user")
		return "", err
	}

	token, err := jwt.CreateToken(&jwt.Payload{
		ID: user.ID,
		Username: user.Username,
		Email: user.Email,
		SecretKey: s.config.JWTKey,
	})
	if err != nil {
		log.Error().Err(err).Msg("Failed to create token")
		return "", err
	}

	return token, nil
}