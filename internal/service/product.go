package service

import (
	"context"
	"log/slog"
	"pvz/internal/model"
)

type ProductRepository interface {
	CreateProduct(ctx context.Context, product model.Product) (model.Product, error)
}

type ProductService struct {
	productRepo   ProductRepository
	receptionRepo ReceptionRepository
	logger        *slog.Logger
}

func NewProductService(pr ProductRepository, rr ReceptionRepository, logger *slog.Logger) *ProductService {
	return &ProductService{productRepo: pr, receptionRepo: rr, logger: logger}
}

func (s *ProductService) CreateProduct(ctx context.Context, product model.Product, pvz model.Pvz) (model.Product, error) {
	reception, err := s.receptionRepo.GetInProgressReception(ctx, pvz.Id)
	if err != nil {
		s.logger.Error("Failed to get in progress reception", "err", err)
		return model.Product{}, err
	}

	product.ReceptionId = reception.Id

	product, err = s.productRepo.CreateProduct(ctx, product)
	if err != nil {
		s.logger.Error("Failed to create product", "err", err)
		return model.Product{}, err
	}

	return product, nil
}
