package main

import (
	"log/slog"
	"pvz/internal/config"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()
	// conn := db.NewPostgresConn(cfg.Postgres)
	// defer conn.Close()

	// userRepo := repository.NewUserRepository(conn)
	// userService := service.NewUserService(userRepo)
	// userHandler := handler.NewUserHandler(userService)

	r := gin.Default()

	slog.Info("Starting server at", "port", cfg.Port)

	r.Run(":" + cfg.Port)
}
