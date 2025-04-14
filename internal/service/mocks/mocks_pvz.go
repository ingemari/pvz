package mocks

import (
	"context"
	"pvz/internal/model"

	"github.com/stretchr/testify/mock"
)

type PvzRepository struct {
	mock.Mock
}

func (m *PvzRepository) CreatePvz(ctx context.Context, pvz model.Pvz) (model.Pvz, error) {
	args := m.Called(ctx, pvz)
	return args.Get(0).(model.Pvz), args.Error(1)
}

func (m *PvzRepository) IsPvz(ctx context.Context, pvz model.Pvz) (bool, error) {
	args := m.Called(ctx, pvz)
	return args.Bool(0), args.Error(1)
}
