package repository

import (
	"context"
	"log/slog"
	"pvz/internal/mapper"
	"pvz/internal/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PvzRepository struct {
	db     *pgxpool.Pool
	logger *slog.Logger
}

func NewPvzRepository(db *pgxpool.Pool, logger *slog.Logger) *PvzRepository {
	return &PvzRepository{db: db, logger: logger}
}

func (r *PvzRepository) CreatePvz(ctx context.Context, pvz model.Pvz) (model.Pvz, error) {
	ent := mapper.PvzToEntity(pvz)

	query := `
		INSERT INTO pvz (city)
		VALUES ($1)
		RETURNING id, registration_date
	`

	err := r.db.QueryRow(ctx, query, ent.City).Scan(&ent.Id, &ent.RegistrationDate)
	if err != nil {
		r.logger.Error("Failed to create PVZ")
		return model.Pvz{}, err
	}

	pvz = mapper.EntityToPvz(ent)
	return pvz, nil
}
