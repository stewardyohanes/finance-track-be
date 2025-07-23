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
	
	tokenPair, err := h.usersService.SignIn(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, users.SignInResponse{
		Data: struct {
			AccessToken  string `json:"access_token"`
			RefreshToken string `json:"refresh_token"`
		}{
			AccessToken:  tokenPair.AccessToken,
			RefreshToken: tokenPair.RefreshToken,
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
	
	tokenPair, user, err := h.usersService.SignUp(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, users.SignUpResponse{
		Data: struct {
			ID           uuid.UUID `json:"id"`
			Username     string    `json:"username"`
			Email        string    `json:"email"`
			CreatedAt    time.Time `json:"created_at"`
			UpdatedAt    time.Time `json:"updated_at"`
			AccessToken  string    `json:"access_token"`
			RefreshToken string    `json:"refresh_token"`
		}{
			ID:           user.ID,
			Username:     user.Username,
			Email:        user.Email,
			CreatedAt:    user.CreatedAt,
			UpdatedAt:    user.UpdatedAt,
			AccessToken:  tokenPair.AccessToken,
			RefreshToken: tokenPair.RefreshToken,
		},
		Message: "Sign up successful",
	})
}

func (h *handler) RefreshToken(c *gin.Context) {
	var req users.RefreshTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	tokenPair, err := h.usersService.RefreshToken(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, users.RefreshTokenResponse{
		Data: struct {
			AccessToken  string `json:"access_token"`
			RefreshToken string `json:"refresh_token"`
		}{
			AccessToken:  tokenPair.AccessToken,
			RefreshToken: tokenPair.RefreshToken,
		},
		Message: "Token refreshed successfully",
	})
}