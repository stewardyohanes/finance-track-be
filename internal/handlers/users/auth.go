package users

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stewardyohanes/finance-tracker/internal/models/users"
)

func (h *handler) SignIn(c *gin.Context) {
	var req users.SignInRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	token, err := h.usersService.SignIn(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, users.SignInResponse{
		Data: struct {
			Token string `json:"token"`
		}{
			Token: token,
		},
		Message: "Sign in successful",
	})
}

func (h *handler) SignUp(c *gin.Context) {
	var req users.SignUpRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	
	token, err := h.usersService.SignUp(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, users.SignUpResponse{
		Data: struct {
			ID        uuid.UUID `json:"id"`
			Username  string    `json:"username"`
			Email     string    `json:"email"`
			CreatedAt time.Time `json:"created_at"`
			UpdatedAt time.Time `json:"updated_at"`
			Token     string    `json:"token"`
		}{
			ID:        uuid.New(),
			Username:  req.Username,
			Email:     req.Email,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Token:     token,
		},
		Message: "Sign up successful",
	})
}