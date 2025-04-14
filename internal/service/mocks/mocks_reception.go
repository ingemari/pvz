package mocks

import (
	"context"
	"pvz/internal/model"

	"github.com/stretchr/testify/mock"
)

type ReceptionRepository struct {
	mock.Mock
}

func (m *ReceptionRepository) GetStatus(ctx context.Context, reception model.Reception) (string, error) {
	args := m.Called(ctx, reception)
	return args.String(0), args.Error(1)
}

func (m *ReceptionRepository) CreateReception(ctx context.Context, reception model.Reception) (model.Reception, error) {
	args := m.Called(ctx, reception)
	return args.Get(0).(model.Reception), args.Error(1)
}
