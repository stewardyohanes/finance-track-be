package users

import (
	"context"
	"errors"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/stewardyohanes/finance-tracker/internal/models/users"
	"github.com/stewardyohanes/finance-tracker/pkg/jwt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func (s *service) SignIn(ctx context.Context, req *users.SignInRequest) (*jwt.TokenPair, error) {
	// Check if the user exists
	user, err := s.usersRepository.GetUserByUsername(req.Username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		log.Error().Err(err).Msg("Failed to get user")
		return nil, err
	}
	
	// Check if the password is correct
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		log.Error().Err(err).Msg("Failed to compare password")
		return nil, errors.New("invalid password")
	}

	// Create token pair
	tokenPair, err := jwt.CreateTokenPair(&jwt.Payload{
		ID: user.ID,
		Username: user.Username,
		Email: user.Email,
		SecretKey: s.config.JWTKey,
	})
	if err != nil {
		log.Error().Err(err).Msg("Failed to create token pair")
		return nil, err
	}

	// Save refresh token to database
	refreshTokenExpiry := time.Now().Add(time.Hour * 24 * 7) // 7 hari
	err = s.usersRepository.UpdateRefreshToken(user.ID, tokenPair.RefreshToken, refreshTokenExpiry)
	if err != nil {
		log.Error().Err(err).Msg("Failed to save refresh token")
		return nil, err
	}

	return tokenPair, nil
}

func (s *service) SignUp(ctx context.Context, req *users.SignUpRequest) (*jwt.TokenPair, *users.User, error) {
	// Check if the user already exists
	existingEmail, err := s.usersRepository.GetUserByEmail(req.Email)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get user by email")
		return nil, nil, err
	}

	existingUsername, err := s.usersRepository.GetUserByUsername(req.Username)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get user by username")
		return nil, nil, err
	}

	if existingEmail != nil {
		return nil, nil, errors.New("email already exists")
	}

	if existingUsername != nil {	
		return nil, nil, errors.New("username already exists")
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Error().Err(err).Msg("Failed to hash password")
		return nil, nil, err
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
		return nil, nil, err
	}

	// Create token pair
	tokenPair, err := jwt.CreateTokenPair(&jwt.Payload{
		ID: user.ID,
		Username: user.Username,
		Email: user.Email,
		SecretKey: s.config.JWTKey,
	})
	if err != nil {
		log.Error().Err(err).Msg("Failed to create token pair")
		return nil, nil, err
	}

	// Save refresh token to database
	refreshTokenExpiry := time.Now().Add(time.Hour * 24 * 7) // 7 hari
	err = s.usersRepository.UpdateRefreshToken(user.ID, tokenPair.RefreshToken, refreshTokenExpiry)
	if err != nil {
		log.Error().Err(err).Msg("Failed to save refresh token")
		return nil, nil, err
	}

	return tokenPair, user, nil
}

func (s *service) RefreshToken(ctx context.Context, req *users.RefreshTokenRequest) (*jwt.TokenPair, error) {
	// Validasi refresh token
	user, err := s.usersRepository.GetUserByRefreshToken(req.RefreshToken)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get user by refresh token")
		return nil, err
	}

	if user == nil {
		return nil, errors.New("invalid or expired refresh token")
	}

	// Clear old refresh token (one-time use)
	err = s.usersRepository.ClearRefreshToken(user.ID)
	if err != nil {
		log.Error().Err(err).Msg("Failed to clear old refresh token")
		return nil, err
	}

	// Create new token pair
	tokenPair, err := jwt.CreateTokenPair(&jwt.Payload{
		ID: user.ID,
		Username: user.Username,
		Email: user.Email,
		SecretKey: s.config.JWTKey,
	})
	if err != nil {
		log.Error().Err(err).Msg("Failed to create new token pair")
		return nil, err
	}

	// Save new refresh token to database
	refreshTokenExpiry := time.Now().Add(time.Hour * 24 * 7) // 7 hari
	err = s.usersRepository.UpdateRefreshToken(user.ID, tokenPair.RefreshToken, refreshTokenExpiry)
	if err != nil {
		log.Error().Err(err).Msg("Failed to save new refresh token")
		return nil, err
	}

	return tokenPair, nil
}