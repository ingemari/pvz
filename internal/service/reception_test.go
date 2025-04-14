package service_test

import (
	"context"
	"errors"
	"testing"

	"pvz/internal/model"
	"pvz/internal/service"
	"pvz/internal/service/mocks"

	"log/slog"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestReceptionService_CreateReception_Success(t *testing.T) {
	pvz := model.Pvz{Id: uuid.New(), City: "Казань"}
	reception := model.Reception{PvzId: pvz.Id}

	mockPvzRepo := new(mocks.PvzRepository)
	mockReceptionRepo := new(mocks.ReceptionRepository)
	logger := slog.Default()

	service := service.NewReceptionService(mockReceptionRepo, mockPvzRepo, logger)

	mockPvzRepo.On("IsPvz", mock.Anything, pvz).Return(true, nil)
	mockReceptionRepo.On("GetStatus", mock.Anything, reception).Return("close", nil)
	mockReceptionRepo.On("CreateReception", mock.Anything, reception).Return(reception, nil)

	result, err := service.CreateReception(context.Background(), reception, pvz)

	assert.NoError(t, err)
	assert.Equal(t, reception, result)
	mockPvzRepo.AssertExpectations(t)
	mockReceptionRepo.AssertExpectations(t)
}

func TestReceptionService_CreateReception_PvzNotExist(t *testing.T) {
	pvz := model.Pvz{Id: uuid.New()}
	reception := model.Reception{PvzId: pvz.Id}

	mockPvzRepo := new(mocks.PvzRepository)
	mockReceptionRepo := new(mocks.ReceptionRepository)
	logger := slog.Default()

	service := service.NewReceptionService(mockReceptionRepo, mockPvzRepo, logger)

	mockPvzRepo.On("IsPvz", mock.Anything, pvz).Return(false, nil)

	_, err := service.CreateReception(context.Background(), reception, pvz)

	assert.ErrorContains(t, err, "incorrect pvz_id")
	mockPvzRepo.AssertExpectations(t)
}

func TestReceptionService_CreateReception_InProgress(t *testing.T) {
	pvz := model.Pvz{Id: uuid.New()}
	reception := model.Reception{PvzId: pvz.Id}

	mockPvzRepo := new(mocks.PvzRepository)
	mockReceptionRepo := new(mocks.ReceptionRepository)
	logger := slog.Default()

	service := service.NewReceptionService(mockReceptionRepo, mockPvzRepo, logger)

	mockPvzRepo.On("IsPvz", mock.Anything, pvz).Return(true, nil)
	mockReceptionRepo.On("GetStatus", mock.Anything, reception).Return("in_progress", nil)

	_, err := service.CreateReception(context.Background(), reception, pvz)

	assert.ErrorContains(t, err, "status is in progress")
}

func TestReceptionService_CreateReception_GetStatusError(t *testing.T) {
	pvz := model.Pvz{Id: uuid.New()}
	reception := model.Reception{PvzId: pvz.Id}

	mockPvzRepo := new(mocks.PvzRepository)
	mockReceptionRepo := new(mocks.ReceptionRepository)
	logger := slog.Default()

	service := service.NewReceptionService(mockReceptionRepo, mockPvzRepo, logger)

	mockPvzRepo.On("IsPvz", mock.Anything, pvz).Return(true, nil)
	mockReceptionRepo.On("GetStatus", mock.Anything, reception).Return("", errors.New("db fail"))

	_, err := service.CreateReception(context.Background(), reception, pvz)

	assert.ErrorContains(t, err, "db fail")
}

func TestReceptionService_CreateReception_CreateError(t *testing.T) {
	pvz := model.Pvz{Id: uuid.New()}
	reception := model.Reception{PvzId: pvz.Id}

	mockPvzRepo := new(mocks.PvzRepository)
	mockReceptionRepo := new(mocks.ReceptionRepository)
	logger := slog.Default()

	service := service.NewReceptionService(mockReceptionRepo, mockPvzRepo, logger)

	mockPvzRepo.On("IsPvz", mock.Anything, pvz).Return(true, nil)
	mockReceptionRepo.On("GetStatus", mock.Anything, reception).Return("close", nil)
	mockReceptionRepo.On("CreateReception", mock.Anything, reception).Return(model.Reception{}, errors.New("insert fail"))

	_, err := service.CreateReception(context.Background(), reception, pvz)

	assert.ErrorContains(t, err, "insert fail")
}
