package router

import (
	"log/slog"
	"pvz/internal/handler"
	"pvz/internal/repository"
	"pvz/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func SetupRouter(db *pgxpool.Pool, logger *slog.Logger) *gin.Engine {
	router := gin.Default()

	userRepo := repository.NewUserRepository(db)
	authService := service.NewAuthService(userRepo)
	authHandler := handler.NewAuthHandler(authService)

	router.POST("/register", authHandler.HandleRegister)
	router.POST("/dummyLogin", authHandler.HandleDummy)
	router.POST("/login", authHandler.HandleLogin)

	logger.Info("Create user endpoint registered")
	return router
}
