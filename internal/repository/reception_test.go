package repository_test

import (
	"context"
	"log/slog"
	"pvz/internal/model"
	"pvz/internal/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateReception_Success(t *testing.T) {
	repo := repository.NewReceptionRepository(db, slog.Default())

	pvzRepo := repository.NewPvzRepository(db, slog.Default())
	pvz, _ := pvzRepo.CreatePvz(context.Background(), model.Pvz{City: "Казань"})

	rec := model.Reception{PvzId: pvz.Id}

	created, err := repo.CreateReception(context.Background(), rec)

	assert.NoError(t, err)
	assert.NotEmpty(t, created.Id)
	assert.Equal(t, "in_progress", created.Status)
}

func TestGetStatus_WithReception(t *testing.T) {
	repo := repository.NewReceptionRepository(db, slog.Default())

	pvzRepo := repository.NewPvzRepository(db, slog.Default())
	pvz, _ := pvzRepo.CreatePvz(context.Background(), model.Pvz{City: "Москва"})

	rec := model.Reception{PvzId: pvz.Id}
	_, _ = repo.CreateReception(context.Background(), rec)

	status, err := repo.GetStatus(context.Background(), rec)

	assert.NoError(t, err)
	assert.Equal(t, "in_progress", status)
}
