package main

import (
	"log"

	"account/internal/config"
	"account/internal/logger"
	"account/internal/repository"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	logger := logger.New(cfg)

	db, err := gorm.Open(postgres.Open(cfg.DbDsn), &gorm.Config{})
	if err != nil {
		logger.Error().Msgf("Failed to connect to database: %v", err)
		return
	}
	logger.Info().Msg("Database connected")

	repo := repository.NewRepository(db, &logger)
	_ = repo

	logger.Info().Msg("Service starting up")
}
