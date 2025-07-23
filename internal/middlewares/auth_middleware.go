package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/stewardyohanes/finance-tracker/config"
	"github.com/stewardyohanes/finance-tracker/pkg/jwt"
)

func AuthMiddleware() gin.HandlerFunc {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load config")
	}
	
	secretKey := cfg.JWTKey

	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		
		header := strings.TrimPrefix(token, "Bearer ")
		if header == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		userID, username, email, err := jwt.ValidateToken(header, secretKey)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		c.Set("userID", userID)
		c.Set("username", username)
		c.Set("email", email)
		c.Next()
	}
}