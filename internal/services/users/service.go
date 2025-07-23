package users

import (
	"time"

	"github.com/google/uuid"
	"github.com/stewardyohanes/finance-tracker/config"
	"github.com/stewardyohanes/finance-tracker/internal/models/users"
)

type usersRepository interface {
	CreateUser(user *users.User) error
	GetUserByEmail(email string) (*users.User, error)
	GetUserByUsername(username string) (*users.User, error)
	GetUserByID(id uuid.UUID) (*users.User, error)
	UpdateRefreshToken(userID uuid.UUID, refreshToken string, expiry time.Time) error
	GetUserByRefreshToken(refreshToken string) (*users.User, error)
	ClearRefreshToken(userID uuid.UUID) error
}

type service struct {
	usersRepository usersRepository
	config          *config.Config
}

func NewService(usersRepository usersRepository, config *config.Config) *service {
	return &service{usersRepository: usersRepository, config: config}
}