package repository

import (
	"context"
	"pvz/internal/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(ctx context.Context, user model.User) (string, error) {
	query := `INSERT INTO users (email, password, role) VALUES ($1, $2, $3) RETURNING id`
	var id string
	err := r.db.QueryRow(ctx, query, user.Email, user.Password, user.Role).Scan(&id)
	return id, err
}
