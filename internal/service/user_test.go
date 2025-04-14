package service_test

import (
	"context"
	"log/slog"
	middleware "pvz/internal/middleware/auth"
	"pvz/internal/model"
	"pvz/internal/service"
	"pvz/internal/service/mocks"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAuthService_CreateUser(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	logger := slog.Default()
	authService := service.NewAuthService(mockRepo, logger)

	user := model.User{
		Id:       uuid.New().String(),
		Email:    "test@example.com",
		Password: "pass123",
		Role:     "employee",
	}

	mockRepo.On("CreateUser", mock.Anything, user).Return(user, nil)

	createdUser, err := authService.CreateUser(context.Background(), user)

	assert.NoError(t, err)
	assert.Equal(t, user, createdUser)
	mockRepo.AssertExpectations(t)
}

func TestAuthService_DummyToken(t *testing.T) {
	logger := slog.Default()
	authService := service.NewAuthService(nil, logger)

	token, err := authService.DummyToken(context.Background(), "moderator")

	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}

func TestAuthService_LoginUser_Success(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	logger := slog.Default()
	authService := service.NewAuthService(mockRepo, logger)

	rawPassword := "secret123"
	hashed, _ := middleware.GeneratePasswordHash(rawPassword)

	user := model.User{
		Email:    "test@example.com",
		Password: rawPassword,
	}

	storedUser := model.User{
		Id:       uuid.New().String(),
		Email:    "test@example.com",
		Password: string(hashed),
		Role:     "employee",
	}

	mockRepo.On("FindByEmail", mock.Anything, user.Email).Return(storedUser, nil)

	token, err := authService.LoginUser(context.Background(), user)

	assert.NoError(t, err)
	assert.NotEmpty(t, token)
	mockRepo.AssertExpectations(t)
}

func TestAuthService_LoginUser_WrongPassword(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	logger := slog.Default()
	authService := service.NewAuthService(mockRepo, logger)

	user := model.User{
		Email:    "test@example.com",
		Password: "wrongpass",
	}

	hashed, _ := middleware.GeneratePasswordHash("correctpass")

	storedUser := model.User{
		Id:       uuid.New().String(),
		Email:    "test@example.com",
		Password: string(hashed),
		Role:     "employee",
	}

	mockRepo.On("FindByEmail", mock.Anything, user.Email).Return(storedUser, nil)

	token, err := authService.LoginUser(context.Background(), user)

	assert.Error(t, err)
	assert.Empty(t, token)
}
