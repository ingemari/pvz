package mocks

import (
	"context"
	"pvz/internal/model"

	"github.com/stretchr/testify/mock"
)

type MockAuthService struct {
	mock.Mock
}

func (m *MockAuthService) CreateUser(ctx context.Context, user model.User) (model.User, error) {
	args := m.Called(ctx, user)
	return args.Get(0).(model.User), args.Error(1)
}

func (m *MockAuthService) DummyToken(ctx context.Context, role string) (string, error) {
	args := m.Called(ctx, role)
	return args.String(0), args.Error(1)
}

func (m *MockAuthService) LoginUser(ctx context.Context, user model.User) (string, error) {
	args := m.Called(ctx, user)
	return args.String(0), args.Error(1)
}
