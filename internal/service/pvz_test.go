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

func TestPvzService_CreatePvz_Success(t *testing.T) {
	mockRepo := new(mocks.PvzRepository)
	logger := slog.Default()
	s := service.NewPvzService(mockRepo, logger)

	input := model.Pvz{City: "Москва"}
	expected := input
	expected.Id = uuid.New()

	mockRepo.On("CreatePvz", mock.Anything, input).Return(expected, nil)

	result, err := s.CreatePvz(context.Background(), input)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	mockRepo.AssertExpectations(t)
}

func TestPvzService_CreatePvz_Error(t *testing.T) {
	mockRepo := new(mocks.PvzRepository)
	logger := slog.Default()
	s := service.NewPvzService(mockRepo, logger)

	input := model.Pvz{City: "Казань"}
	mockRepo.On("CreatePvz", mock.Anything, input).Return(model.Pvz{}, errors.New("db error"))

	result, err := s.CreatePvz(context.Background(), input)

	assert.Error(t, err)
	assert.Equal(t, model.Pvz{}, result)
	assert.EqualError(t, err, "db error")
	mockRepo.AssertExpectations(t)
}
