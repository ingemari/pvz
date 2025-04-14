package handler_test

import (
	"bytes"
	"errors"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"pvz/internal/handler"
	"pvz/internal/handler/mocks"
	"pvz/internal/model"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAuthHandler_HandleRegister(t *testing.T) {
	tests := []struct {
		name         string
		requestBody  string
		mockReturn   model.User
		mockError    error
		expectedCode int
		expectedBody string
	}{
		{
			name:        "Successful registration",
			requestBody: `{"email":"test@example.com","password":"ValidPass123!","role":"employee"}`,
			mockReturn: model.User{
				Id:    uuid.New().String(),
				Email: "test@example.com",
				Role:  "employee",
			},
			expectedCode: http.StatusCreated,
			expectedBody: `{"id":".+","email":"test@example.com","role":"employee"}`,
		},
		{
			name:         "Invalid email format",
			requestBody:  `{"email":"invalid","password":"ValidPass123!","role":"employee"}`,
			expectedCode: http.StatusBadRequest,
			expectedBody: `{"message":"Invalid email"}`,
		},
		{
			name:         "Invalid role",
			requestBody:  `{"email":"test@example.com","password":"ValidPass123!","role":"invalid"}`,
			expectedCode: http.StatusBadRequest,
			expectedBody: `{"message":"Invalid role"}`,
		},
		{
			name:         "Weak password",
			requestBody:  `{"email":"test@example.com","password":"short","role":"employee"}`,
			expectedCode: http.StatusBadRequest,
			expectedBody: `{"message":"Invalid pass"}`,
		},
		{
			name:         "User already exists",
			requestBody:  `{"email":"exists@example.com","password":"ValidPass123!","role":"employee"}`,
			mockError:    errors.New("user exists"),
			expectedCode: http.StatusInternalServerError,
			expectedBody: `{"message":"User already exist"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService := new(mocks.MockAuthService)
			logger := slog.Default()
			handler := handler.NewAuthHandler(mockService, logger)

			if tt.mockReturn.Email != "" || tt.mockError != nil {
				mockService.On("CreateUser", mock.Anything, mock.AnythingOfType("model.User")).
					Return(tt.mockReturn, tt.mockError)
			}

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request, _ = http.NewRequest(
				http.MethodPost,
				"/register",
				bytes.NewBuffer([]byte(tt.requestBody)),
			)
			c.Request.Header.Set("Content-Type", "application/json")

			handler.HandleRegister(c)

			assert.Equal(t, tt.expectedCode, w.Code)
			assert.Regexp(t, tt.expectedBody, w.Body.String())

			mockService.AssertExpectations(t)
		})
	}
}

func TestAuthHandler_HandleDummy(t *testing.T) {
	tests := []struct {
		name         string
		requestBody  string
		mockToken    string
		mockError    error
		expectedCode int
		expectedBody string
	}{
		{
			name:         "Success employee role",
			requestBody:  `{"role":"employee"}`,
			mockToken:    "dummy_token_123",
			expectedCode: http.StatusOK,
			expectedBody: `"dummy_token_123"`,
		},
		{
			name:         "Success moderator role",
			requestBody:  `{"role":"moderator"}`,
			mockToken:    "dummy_token_456",
			expectedCode: http.StatusOK,
			expectedBody: `"dummy_token_456"`,
		},
		{
			name:         "Invalid role",
			requestBody:  `{"role":"invalid"}`,
			expectedCode: http.StatusBadRequest,
			expectedBody: `{"message":"Invalid role"}`,
		},
		{
			name:         "Service error",
			requestBody:  `{"role":"employee"}`,
			mockError:    errors.New("service error"),
			expectedCode: http.StatusInternalServerError,
			expectedBody: `{"message":"service error"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService := new(mocks.MockAuthService)
			logger := slog.Default()
			handler := handler.NewAuthHandler(mockService, logger)

			if tt.mockToken != "" || tt.mockError != nil {
				mockService.On("DummyToken", mock.Anything, mock.AnythingOfType("string")).
					Return(tt.mockToken, tt.mockError)
			}

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request, _ = http.NewRequest(
				http.MethodPost,
				"/dummy",
				bytes.NewBuffer([]byte(tt.requestBody)),
			)
			c.Request.Header.Set("Content-Type", "application/json")

			handler.HandleDummy(c)

			assert.Equal(t, tt.expectedCode, w.Code)
			assert.JSONEq(t, tt.expectedBody, w.Body.String())

			mockService.AssertExpectations(t)
		})
	}
}

func TestAuthHandler_HandleLogin(t *testing.T) {
	tests := []struct {
		name         string
		requestBody  string
		mockToken    string
		mockError    error
		expectedCode int
		expectedBody string
	}{
		{
			name:         "Successful login",
			requestBody:  `{"email":"valid@example.com","password":"correctPassword"}`,
			mockToken:    "auth_token_123",
			expectedCode: http.StatusOK,
			expectedBody: `"auth_token_123"`,
		},
		{
			name:         "Invalid email format",
			requestBody:  `{"email":"invalid","password":"password"}`,
			expectedCode: http.StatusBadRequest,
			expectedBody: `{"message":"Invalid email format"}`,
		},
		{
			name:         "Short password",
			requestBody:  `{"email":"valid@example.com","password":"short"}`,
			expectedCode: http.StatusBadRequest,
			expectedBody: `{"message":"Password must be at least 6 characters"}`,
		},
		{
			name:         "Auth service error",
			requestBody:  `{"email":"valid@example.com","password":"correctPassword"}`,
			mockError:    errors.New("auth failed"),
			expectedCode: http.StatusInternalServerError,
			expectedBody: `{"message":"auth failed"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService := new(mocks.MockAuthService)
			logger := slog.Default()
			handler := handler.NewAuthHandler(mockService, logger)

			if tt.mockToken != "" || tt.mockError != nil {
				mockService.On("LoginUser", mock.Anything, mock.AnythingOfType("model.User")).
					Return(tt.mockToken, tt.mockError)
			}

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request, _ = http.NewRequest(
				http.MethodPost,
				"/login",
				bytes.NewBuffer([]byte(tt.requestBody)),
			)
			c.Request.Header.Set("Content-Type", "application/json")

			handler.HandleLogin(c)

			assert.Equal(t, tt.expectedCode, w.Code)
			assert.JSONEq(t, tt.expectedBody, w.Body.String())

			mockService.AssertExpectations(t)
		})
	}
}
