package router

import (
	"log/slog"
	"pvz/internal/handler"
	middleware "pvz/internal/middleware/auth"
	"pvz/internal/repository"
	"pvz/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func SetupRouter(db *pgxpool.Pool, logger *slog.Logger) *gin.Engine {
	router := gin.Default()

	// repo
	userRepo := repository.NewUserRepository(db, logger)
	pvzRepo := repository.NewPvzRepository(db, logger)
	// service
	authService := service.NewAuthService(userRepo, logger)
	pvzService := service.NewPvzService(pvzRepo, logger)
	// handler
	authHandler := handler.NewAuthHandler(authService, logger)
	pvzHandler := handler.NewPvzHandler(pvzService, logger)

	// open routes
	router.POST("/register", authHandler.HandleRegister)
	router.POST("/dummyLogin", authHandler.HandleDummy)
	router.POST("/login", authHandler.HandleLogin)
	// protected group
	protected := router.Group("/")
	protected.Use(middleware.RequireAuth())
	// protecred routes
	protected.POST("/pvz", middleware.RequireRole("moderator"), pvzHandler.HandleCreatePvz)

	logger.Info("Create user endpoint registered")
	return router
}
