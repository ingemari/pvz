package repository

import (
	"context"
	"log/slog"
	"pvz/internal/model"
	"pvz/internal/repository/entities"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	db     *pgxpool.Pool
	logger *slog.Logger
}

func NewUserRepository(db *pgxpool.Pool, logger *slog.Logger) *UserRepository {
	return &UserRepository{db: db, logger: logger}
}

func (r *UserRepository) CreateUser(ctx context.Context, user model.User) (model.User, error) {
	ent := entities.UserToEntity(user)

	query := `
		INSERT INTO users (email, password, role)
		VALUES ($1, $2, $3)
		RETURNING id
	`
	err := r.db.QueryRow(ctx, query, ent.Email, ent.Password, ent.Role).Scan(&ent.Id)
	if err != nil {
		r.logger.Error("Failed to create user (already exist)")
		return model.User{}, err
	}

	createdUser := entities.EntityToUser(ent)
	return createdUser, nil
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (model.User, error) {
	query := `
		SELECT id, email, password, role 
		FROM users 
		WHERE email = $1
	`

	var ent entities.User

	err := r.db.QueryRow(ctx, query, email).Scan(&ent.Id, &ent.Email, &ent.Password, &ent.Role)
	if err != nil {
		r.logger.Error("Failed to find user", "email", email)
		return model.User{}, err
	}

	return entities.EntityToUser(ent), nil
}
