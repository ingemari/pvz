package service

import (
	"context"
	"fmt"

	middleware "pvz/internal/middleware/auth"
	"pvz/internal/model"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user model.User) (model.User, error)
	FindByEmail(ctx context.Context, email string) (model.User, error)
}

type AuthService struct {
	userRepo UserRepository
}

func NewAuthService(u UserRepository) *AuthService {
	return &AuthService{userRepo: u}
}

func (s *AuthService) CreateUser(ctx context.Context, user model.User) (model.User, error) {
	return s.userRepo.CreateUser(ctx, user)
}

func (s *AuthService) DummyToken(ctx context.Context, role string) (string, error) {
	token, err := middleware.GenerateJWT(role)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s *AuthService) LoginUser(ctx context.Context, user model.User) (string, error) {
	// 1. Найти пользователя по email
	storedUser, err := s.userRepo.FindByEmail(ctx, user.Email)
	if err != nil {
		return "", fmt.Errorf("invalid email or password") // обобщённо, чтобы не палить детали
	}

	// 2. Проверить пароль через bcrypt
	if !middleware.CheckPasswordHash(user.Password, storedUser.Password) {
		return "", fmt.Errorf("invalid email or password")
	}

	// 3. Генерация токена по роли
	token, err := middleware.GenerateJWT(storedUser.Role)
	if err != nil {
		return "", err
	}

	return token, nil
}
