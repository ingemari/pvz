package handler_test

import (
	"bytes"
	"encoding/json"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"pvz/internal/handler"
	"pvz/internal/handler/dto"
	"pvz/internal/handler/mocks"
	"pvz/internal/model"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestHandleReceptionCreate_Success(t *testing.T) {
	// Инициализация
	mockService := new(mocks.MockReceptionService)
	logger := slog.Default()
	h := handler.NewReceptionHandler(mockService, logger)

	// Тестовые данные
	testPvzID := uuid.New().String()
	testReq := dto.ReceptionRequst{
		PvzID: testPvzID,
	}

	expectedReception := model.Reception{
		Id:       uuid.New(),
		PvzId:    uuid.MustParse(testPvzID),
		DateTime: time.Now(),
		Status:   "active",
	}

	// Настройка мока
	mockService.On("CreateReception",
		mock.Anything,
		mock.MatchedBy(func(r model.Reception) bool {
			return r.PvzId.String() == testPvzID
		}),
		mock.AnythingOfType("model.Pvz"),
	).Return(expectedReception, nil)

	// Создание тестового запроса
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	reqBody, _ := json.Marshal(testReq)
	c.Request, _ = http.NewRequest(http.MethodPost, "/receptions", bytes.NewBuffer(reqBody))
	c.Request.Header.Set("Content-Type", "application/json")

	// Вызов обработчика
	h.HandleReceptionCreate(c)

	// Проверки
	assert.Equal(t, http.StatusCreated, w.Code)

	var resp dto.ReceptionResponse
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	assert.NoError(t, err)

	assert.Equal(t, expectedReception.Id.String(), resp.Id)
	assert.Equal(t, expectedReception.PvzId.String(), resp.PvzID)
	assert.Equal(t, expectedReception.Status, resp.Status)
	assert.NotEmpty(t, resp.DateTime) // Проверяем что дата не пустая

	mockService.AssertExpectations(t)
}
