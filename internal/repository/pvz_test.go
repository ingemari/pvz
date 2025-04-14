package repository_test

import (
	"context"
	"os"
	"testing"

	"log/slog"
	"pvz/internal/model"
	"pvz/internal/repository"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

var db *pgxpool.Pool

func TestMain(m *testing.M) {
	_ = godotenv.Load(".env.test")

	dbUrl := os.Getenv("TEST_DB_URL")
	pool, err := pgxpool.New(context.Background(), dbUrl)
	if err != nil {
		panic(err)
	}
	db = pool
	code := m.Run()
	pool.Close()
	os.Exit(code)
}

func TestCreatePvz_Success(t *testing.T) {
	repo := repository.NewPvzRepository(db, slog.Default())

	input := model.Pvz{City: "Казань"}

	result, err := repo.CreatePvz(context.Background(), input)

	assert.NoError(t, err)
	assert.Equal(t, input.City, result.City)
	assert.NotEmpty(t, result.Id)
	assert.False(t, result.RegistrationDate.IsZero())
}

func TestIsPvz_Found(t *testing.T) {
	repo := repository.NewPvzRepository(db, slog.Default())

	created, _ := repo.CreatePvz(context.Background(), model.Pvz{City: "Москва"})

	found, err := repo.IsPvz(context.Background(), created)

	assert.NoError(t, err)
	assert.True(t, found)
}

func TestIsPvz_NotFound(t *testing.T) {
	repo := repository.NewPvzRepository(db, slog.Default())

	notExist := model.Pvz{Id: uuid.New()} // случайный ID

	found, err := repo.IsPvz(context.Background(), notExist)

	assert.NoError(t, err)
	assert.False(t, found)
}
