package users

import (
	"github.com/stewardyohanes/finance-tracker/config"
	"github.com/stewardyohanes/finance-tracker/internal/models/users"
)

type usersRepository interface {
	CreateUser(user *users.User) error
	GetUserByEmail(email string) (*users.User, error)
	GetUserByUsername(username string) (*users.User, error)
}

type service struct {
	usersRepository usersRepository
	config          *config.Config
}

func NewService(usersRepository usersRepository, config *config.Config) *service {
	return &service{usersRepository: usersRepository, config: config}
}