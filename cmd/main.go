package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/stewardyohanes/finance-tracker/config"
	usersHandler "github.com/stewardyohanes/finance-tracker/internal/handlers/users"
	"github.com/stewardyohanes/finance-tracker/internal/models/users"
	usersRepo "github.com/stewardyohanes/finance-tracker/internal/repositories/users"
	"github.com/stewardyohanes/finance-tracker/internal/routes"
	usersService "github.com/stewardyohanes/finance-tracker/internal/services/users"
	"github.com/stewardyohanes/finance-tracker/pkg/database"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load config")
	}

	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to database")
	}

	db.AutoMigrate(&users.User{})

	router := gin.Default()
	routes.SetupRoutes(router)

	usersRepo := usersRepo.NewRepository(db)
	usersService := usersService.NewService(usersRepo, config)
	usersHandler := usersHandler.NewHandler(router, usersService)

	usersHandler.AuthRoutes()

	router.Run(":8080")
}