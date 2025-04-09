package main

import (
	"log/slog"
	"pvz/internal/config"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()
	// db := db.NewPool(cfg)

	// userRepo := repository.NewUserRepository(db)
	// userService := service.NewUserService(userRepo)
	// apiHandler := api.NewHandler(userService)

	r := gin.Default()
	slog.Info("Starting server at", "port", cfg.ServerPort)
	r.Run(":" + cfg.ServerPort)
}

// req
// {
//   "role": "employee"
// }

// resp "string"
