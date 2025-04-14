package service

import (
	"context"
	"errors"
	"log/slog"
	"pvz/internal/model"
)

type ReceptionRepository interface {
	GetStatus(context.Context, model.Reception) (string, error)
	CreateReception(ctx context.Context, reception model.Reception) (model.Reception, error)
}

type ReceptionService struct {
	receptionRepo ReceptionRepository
	pvzRepo       PvzRepository
	logger        *slog.Logger
}

func NewReceptionService(r ReceptionRepository, p PvzRepository, logger *slog.Logger) *ReceptionService {
	return &ReceptionService{receptionRepo: r, pvzRepo: p, logger: logger}
}

func (s *ReceptionService) CreateReception(ctx context.Context, reception model.Reception, pvz model.Pvz) (model.Reception, error) {
	ok, err := s.pvzRepo.IsPvz(ctx, pvz)
	if err != nil {
		s.logger.Error("Error checking pvz", "error", err)
		return model.Reception{}, err
	}
	if !ok {
		s.logger.Error("Error have not pvz")
		return model.Reception{}, errors.New("incorrect pvz_id")
	}
	status, err := s.receptionRepo.GetStatus(ctx, reception)
	if err != nil {
		return model.Reception{}, err
	}
	if status == "in_progress" {
		s.logger.Error("Error create reception! status is in progress")
		return model.Reception{}, errors.New("status is in progress")
	}
	m, err := s.receptionRepo.CreateReception(ctx, reception)
	if err != nil {
		s.logger.Error("Error create reception")
		return model.Reception{}, err
	}
	return m, nil
}
