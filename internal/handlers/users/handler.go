package users

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/stewardyohanes/finance-tracker/internal/models/users"
	"github.com/stewardyohanes/finance-tracker/pkg/jwt"
)

type usersService interface {
	SignIn(ctx context.Context, req *users.SignInRequest) (*jwt.TokenPair, error)
	SignUp(ctx context.Context, req *users.SignUpRequest) (*jwt.TokenPair, *users.User, error)
	RefreshToken(ctx context.Context, req *users.RefreshTokenRequest) (*jwt.TokenPair, error)
}

type handler struct {
	*gin.Engine

	usersService usersService
}

func NewHandler(api *gin.Engine, usersService usersService) *handler {
	return &handler{Engine: api, usersService: usersService}
}

func (h *handler) AuthRoutes() {
	auth := h.Engine.Group("/api/v1")

	auth.POST("/auth/signin", h.SignIn)
	auth.POST("/auth/signup", h.SignUp)
	auth.POST("/auth/refresh", h.RefreshToken)
}

