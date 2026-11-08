package main

import (
	"fmt"
	"log"
	"net"

	_ "account/docs"
	"account/internal/config"
	"account/internal/logger"
	"account/internal/repository"

	//	"github.com/gin-gonic/gin"
	// swaggerFiles "github.com/swaggo/files"
	// ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// @title Account Service
// @version 1.0
// @description Account Service in sample project
// @Host localhost:8080
// @BasePath /
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

	listenAddr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	lis, err := net.Listen("tcp", listenAddr)
	if err != nil {
		logger.Error().Msgf("Failed to listen on %s: %v", listenAddr, err)
		return
	}

	grpcServer := grpc.NewServer()

	// Registration of health service and reflection for debug

	healthSrv := health.NewServer()
	grpc_health_v1.RegisterHealthServer(grpcServer, healthSrv)
	reflection.Register(grpcServer)

	logger.Info().Msgf("gRPC server listening on %s", listenAddr)
	if err := grpcServer.Serve(lis); err != nil {
		logger.Error().Msgf("Failed to serve gRPC: %v", err)
		return
	}

	router := gin.Default()
	err = router.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		logger.Error().Msgf("Failed to set trusted proxies: %v", err)
		return
	}

	logger.Info().Msg("Service starting up")
}
