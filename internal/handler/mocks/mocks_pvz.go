package mocks

import (
	"context"
	"pvz/internal/model"

	"github.com/stretchr/testify/mock"
)

// MockPvzService - это мок для сервиса Pvz
type MockPvzService struct {
	mock.Mock
}

// CreatePvz - мокаем метод CreatePvz
func (m *MockPvzService) CreatePvz(ctx context.Context, pvz model.Pvz) (model.Pvz, error) {
	args := m.Called(ctx, pvz)
	return args.Get(0).(model.Pvz), args.Error(1)
}
