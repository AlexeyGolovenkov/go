package main

import (
	"account/internal/app"
	"account/internal/config"
	"account/internal/logger"
	"context"
)

// @title Account Service
// @version 1.0
// @description Account Service in sample project
// @Host localhost:8080
// @BasePath /
func main() {
	ctx := context.Background()
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	logger := logger.New(cfg)

	application := app.New(&logger, cfg)

	if err := application.Run(ctx); err != nil {
		logger.Fatal().Err(err).Msg("error")
	}
}
