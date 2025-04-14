package repository

import (
	"context"
	"log/slog"
	"pvz/internal/model"
	"pvz/internal/repository/entities"

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
	ent := entities.PvzToEntity(pvz)

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

	pvz = entities.EntityToPvz(ent)
	return pvz, nil
}

// ну нужен тк можно положиться на референс sql
func (r *PvzRepository) IsPvz(ctx context.Context, pvz model.Pvz) (bool, error) {
	id := pvz.Id // ENTITTY!!!

	query := `
		SELECT id
		FROM pvz
		WHERE id = $1
	`
	row := r.db.QueryRow(ctx, query, id)
	if row != nil {
		return true, nil
	}
	return false, nil
}
