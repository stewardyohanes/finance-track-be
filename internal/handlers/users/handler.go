package users

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/stewardyohanes/finance-tracker/internal/models/users"
)

type usersService interface {
	SignIn(ctx context.Context, req *users.SignInRequest) (string, error)
	SignUp(ctx context.Context, req *users.SignUpRequest) (string, error)
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
}

