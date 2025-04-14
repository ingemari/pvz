// handler/mocks/reception_service.go
package mocks

import (
	"context"
	"pvz/internal/model"

	"github.com/stretchr/testify/mock"
)

type MockReceptionService struct {
	mock.Mock
}

func (m *MockReceptionService) CreateReception(ctx context.Context, reception model.Reception, pvz model.Pvz) (model.Reception, error) {
	args := m.Called(ctx, reception, pvz)
	return args.Get(0).(model.Reception), args.Error(1)
}
