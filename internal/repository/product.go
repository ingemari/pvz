package repository

import (
	"context"
	"log/slog"
	"pvz/internal/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ProductRepository struct {
	db     *pgxpool.Pool
	logger *slog.Logger
}

func NewProductRepository(db *pgxpool.Pool, logger *slog.Logger) *ProductRepository {
	return &ProductRepository{db: db, logger: logger}
}

func (r *ProductRepository) CreateProduct(ctx context.Context, product model.Product) (model.Product, error) {
	query := `
		INSERT INTO products (type, reception_id)
		VALUES ($1, $2)
		RETURNING id, date_time
	`

	err := r.db.QueryRow(ctx, query, product.Type, product.ReceptionId).Scan(&product.Id, &product.DateTime)
	if err != nil {
		r.logger.Error("Failed to insert product")
		return model.Product{}, err
	}

	return product, nil
}
