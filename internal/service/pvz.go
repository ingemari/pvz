package service

import (
	"context"
	"log/slog"
	"pvz/internal/model"
)

type PvzRepository interface {
	CreatePvz(ctx context.Context, pvz model.Pvz) (model.Pvz, error)
	IsPvz(ctx context.Context, pvz model.Pvz) (bool, error)
}

type PvzService struct {
	pvzRepo PvzRepository
	logger  *slog.Logger
}

func NewPvzService(pr PvzRepository, logger *slog.Logger) *PvzService {
	return &PvzService{pvzRepo: pr, logger: logger}
}

func (s *PvzService) CreatePvz(ctx context.Context, pvz model.Pvz) (model.Pvz, error) {
	return s.pvzRepo.CreatePvz(ctx, pvz)
}
