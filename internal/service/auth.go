package service

import (
	"context"
	"log/slog"
	"pvz/internal/middleware"
	"pvz/internal/middleware/mapper"
	"pvz/internal/model"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user model.User) (string, error)
}

type AuthService struct {
	userRepo UserRepository
}

func NewAuthService(u UserRepository) *AuthService {
	return &AuthService{userRepo: u}
}

func (s *AuthService) Register(ctx context.Context, req model.RegisterRequest) (*model.RegisterResponse, error) {
	// Хэшируем пароль
	hashedPassword, err := middleware.GeneratePasswordHash(req.Password)
	if err != nil {
		slog.Warn("failed to hash password", "error", err)
		return nil, err
	}

	// Преобразуем RegisterRequest в User (используем маппер)
	user := mapper.RegisterRequestToUser(req, hashedPassword) // передаём []byte

	// Сохраняем в БД
	id, err := s.userRepo.CreateUser(ctx, user)
	if err != nil {
		slog.Error("failed to create user", "error", err)
		return nil, err
	}

	user.ID = id
	// Преобразуем User в RegisterResponse (используем маппер)
	resp := mapper.UserToRegisterResponse(user)
	return &resp, nil
}
