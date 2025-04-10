package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"pvz/internal/config"
	"pvz/internal/db"
	"pvz/internal/handler"
	"pvz/internal/repository"
	"pvz/internal/service"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	// Инициализация логгера с JSON-форматом
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
	slog.SetDefault(logger)

	// Загрузка конфигурации
	cfg := config.LoadConfig()
	logger.Info("Configuration loaded", "config", cfg)

	// Инициализация базы данных
	db := db.InitDB(config.MakeDSN(*cfg))
	defer db.Close()
	logger.Info("Database connection established")

	// Инициализация слоёв приложения
	userRepo := repository.NewUserRepository(db)
	authService := service.NewAuthService(userRepo)
	authHandler := handler.NewAuthHandler(authService)

	// Настройка маршрутизатора
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		c.String(http.StatusOK, "Welcome Gin Server")
	})

	router.POST("/register", authHandler.Register)

	// Настройка HTTP сервера
	srv := &http.Server{
		Addr:    cfg.ServerPort,
		Handler: router,
	}

	// Запуск сервера в отдельной goroutine
	go func() {
		logger.Info("Starting server", "address", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Panicf("Server failed to start %s", err.Error())
		}
	}()

	// Канал для обработки сигналов завершения
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Ожидание сигнала завершения
	sig := <-quit
	logger.Info("Received shutdown signal", "signal", sig)

	// Настройка контекста с таймаутом для graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Graceful shutdown сервера
	logger.Info("Shutting down server...")
	if err := srv.Shutdown(ctx); err != nil {
		log.Panicf("Server shutdown failed %s", err)
	}

	// Проверка завершения shutdown
	select {
	case <-ctx.Done():
		slog.Info("timeout of 5 seconds.")
	}

	logger.Info("Server exited properly")
}
