package main

import (
	"account/internal/config"
	"account/internal/logger"
	"log"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	logger := logger.New(cfg)

	logger.Info().Msg("Service starting up")
}
