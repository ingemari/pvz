package repository_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"pvz/internal/handler"
	"pvz/internal/handler/dto"
	"pvz/internal/handler/mocks"
	"pvz/internal/middleware/logs"
	"pvz/internal/model"
	"pvz/internal/repository"
)

func TestUserRepository_CreateUser_Success(t *testing.T) {
	repo := repository.NewUserRepository(db, slog.Default())

	email := "test_" + uuid.New().String() + "@example.com"
	user := model.User{
		Email:    email,
		Password: "hashed-password",
		Role:     "employee",
	}

	created, err := repo.CreateUser(context.Background(), user)

	assert.NoError(t, err)
	assert.Equal(t, user.Email, created.Email)
	assert.Equal(t, user.Role, created.Role)
	assert.NotEmpty(t, created.Id)
}

func TestHandleCreatePvz_Success(t *testing.T) {
	// Подготовка
	mockPvzService := new(mocks.MockPvzService)
	logger := logs.SetupLogger()
	handler := handler.NewPvzHandler(mockPvzService, logger)

	// Тестовые данные
	pvzID := uuid.New()
	expectedPvz := model.Pvz{
		Id:   pvzID,
		City: "Москва",
		// RegistrationDate намеренно не проверяем
	}

	// Мокаем сервис
	mockPvzService.On("CreatePvz", mock.Anything, mock.MatchedBy(func(p model.Pvz) bool {
		return p.City == "Москва"
	})).Return(expectedPvz, nil)

	// Запрос
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest(
		http.MethodPost,
		"/pvz",
		bytes.NewBuffer([]byte(`{"city":"Москва"}`)),
	)

	// Вызов обработчика
	handler.HandleCreatePvz(c)

	// Проверки
	assert.Equal(t, http.StatusCreated, w.Code)

	var resp dto.PvzCreateResponse
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	assert.NoError(t, err)

	// Проверяем только ID и город
	assert.Equal(t, pvzID.String(), resp.Id)
	assert.Equal(t, "Москва", resp.City)

	// Проверяем, что дата вообще есть (опционально)
	assert.NotEmpty(t, resp.RegDate)

	mockPvzService.AssertExpectations(t)
}
