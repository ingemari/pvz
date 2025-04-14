package repository

import (
	"context"
	"errors"
	"log/slog"
	"pvz/internal/model"
	"pvz/internal/repository/entities"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ReceptionRepository struct {
	db     *pgxpool.Pool
	logger *slog.Logger
}

func NewReceptionRepository(db *pgxpool.Pool, logger *slog.Logger) *ReceptionRepository {
	return &ReceptionRepository{db: db, logger: logger}
}

func (r *ReceptionRepository) GetStatus(ctx context.Context, reception model.Reception) (string, error) {
	ent := entities.ReceptionToEntity(reception)

	query := `
		SELECT status
		FROM receptions
		WHERE pvz_id = $1
		ORDER BY date_time DESC
		LIMIT 1
	`

	err := r.db.QueryRow(ctx, query, ent.PvzId).Scan(&ent.Status)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		r.logger.Error("Failed to create reception STATUS", "pvzID", ent.PvzId)
		return "", err
	}

	return ent.Status, nil
}

func (r *ReceptionRepository) CreateReception(ctx context.Context, reception model.Reception) (model.Reception, error) {
	ent := entities.ReceptionToEntity(reception)
	ent.Status = "in_progress"

	query := `
		INSERT INTO receptions (pvz_id, status)
		VALUES ($1, $2)
		RETURNING id, date_time
	`

	err := r.db.QueryRow(ctx, query, ent.PvzId, ent.Status).Scan(&ent.Id, &ent.DateTime)
	if err != nil {
		r.logger.Error("Failed to create reception CREATE", "pvz_id", ent.PvzId)
		return model.Reception{}, err
	}

	reception = entities.EntityToReception(ent)

	return reception, nil
}

func (r *ReceptionRepository) GetInProgressReception(ctx context.Context, pvzId uuid.UUID) (model.Reception, error) {
	reception := model.Reception{
		PvzId: pvzId,
	}

	query := `
		SELECT id, date_time, status
		FROM receptions
		WHERE pvz_id = $1 AND status = 'in_progress'
		ORDER BY date_time DESC
		LIMIT 1
	`

	err := r.db.QueryRow(ctx, query, pvzId).Scan(&reception.Id, &reception.DateTime, &reception.Status)
	if err != nil {
		r.logger.Error("Failed to find reception in_progress", "pvz_id", pvzId)
		return model.Reception{}, err
	}

	return reception, nil
}
