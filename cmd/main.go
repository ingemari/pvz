package main

import (
	"log/slog"
	"pvz/internal/config"
	"pvz/internal/db"
	"pvz/internal/handler/router"
	"pvz/internal/middleware/logs"
	"pvz/internal/server"
	"time"
)

func main() {
	logger := logs.SetupLogger()
	slog.SetDefault(logger)

	cfg := config.LoadConfig()
	logger.Info("Configuration loaded", "config", cfg)

	db := db.InitDB(config.MakeDSN(*cfg))
	defer db.Close()
	logger.Info("Database connection established")

	router := router.SetupRouter(db, logger)

	server.Run(
		logger,
		router,
		cfg.ServerPort,
		30*time.Second,
	)
}
