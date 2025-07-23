package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/stewardyohanes/finance-tracker/internal/models/users"
	"github.com/stewardyohanes/finance-tracker/internal/routes"
	"github.com/stewardyohanes/finance-tracker/pkg/database"
)

func main() {
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to database")
	}

	db.AutoMigrate(&users.UserModel{})

	router := gin.Default()
	routes.SetupRoutes(router)

	router.Run(":8080")
}