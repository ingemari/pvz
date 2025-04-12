package repository

import (
	"context"
	"log/slog"
	"pvz/internal/mapper"
	"pvz/internal/model"
	"pvz/internal/repository/entities"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(ctx context.Context, user model.User) (model.User, error) {
	ent := mapper.UserToEntity(user)

	query := `
		INSERT INTO users (email, password, role)
		VALUES ($1, $2, $3)
		RETURNING id
	`
	var id uuid.UUID
	err := r.db.QueryRow(ctx, query, ent.Email, ent.Password, ent.Role).Scan(&id)
	if err != nil {
		slog.Error("Failed to create user on storage layer", "err", err)
		return model.User{}, err
	}

	ent.ID = id

	createdUser := mapper.EntityToUser(ent)
	return createdUser, nil
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (model.User, error) {
	query := `SELECT id, email, password, role FROM users WHERE email = $1`
	var ent entities.User

	err := r.db.QueryRow(ctx, query, email).Scan(&ent.ID, &ent.Email, &ent.Password, &ent.Role)
	if err != nil {
		slog.Error("Failed to find user by email", "err", err)
		return model.User{}, err
	}

	return mapper.EntityToUser(ent), nil
}
