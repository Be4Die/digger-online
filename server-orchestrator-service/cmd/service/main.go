package main

import (
	"fmt"
	"log"

	"github.com/Be4Die/digger-online/server-orchestrator-service/internal/api/handlers"
	"github.com/Be4Die/digger-online/server-orchestrator-service/internal/config"
	"github.com/gin-gonic/gin"
)

func main() {
	// Загружаем конфигурацию
	cfg := config.LoadConfig()

	if cfg.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	router := gin.Default()

	healthHandler := handlers.NewHealthHandler()
	serverHandler := handlers.NewServerHandler()

	v1 := router.Group("/api/v1")
	{
		v1.GET("/health", healthHandler.HealthCheck)
		v1.POST("/servers", serverHandler.CreateServer)
		v1.GET("/servers/available", serverHandler.GetAvailableServer)
	}

	addr := fmt.Sprintf(":%d", cfg.APIPort)
	log.Printf("Starting server on %s in %s mode", addr, cfg.Env)
	
	if err := router.Run(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}