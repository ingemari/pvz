package handler

import (
	"context"
	"log/slog"
	"net/http"
	"pvz/internal/handler/dto"
	"pvz/internal/mapper"
	middleware "pvz/internal/middleware/auth"
	"pvz/internal/middleware/validations"
	"pvz/internal/model"

	"github.com/gin-gonic/gin"
)

type AuthService interface {
	CreateUser(ctx context.Context, user model.User) (model.User, error)
	DummyToken(ctx context.Context, role string) (string, error)
	LoginUser(ctx context.Context, user model.User) (string, error)
}

type AuthHandler struct {
	authService AuthService
	logger      *slog.Logger
}

func NewAuthHandler(as AuthService, logger *slog.Logger) *AuthHandler {
	return &AuthHandler{authService: as, logger: logger}
}

func (h *AuthHandler) HandleLogin(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Error("Invalid requst LOGIN")
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: "Invalid request format"})
		return
	}

	// Валидация email
	if !validations.IsValidEmail(req.Email) {
		h.logger.Error("Fail email validations LOGIN")
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: "Invalid email format"})
		return
	}

	// Легкая валидация данных
	if len(req.Password) < 6 {
		h.logger.Error("Fail password validations LOGIN")
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: "Password must be at least 6 characters"})
		return
	}

	// mapping in model
	model := mapper.LoginRequestToUser(req, []byte(req.Password))

	token, err := h.authService.LoginUser(c.Request.Context(), model)
	if err != nil {
		h.logger.Error("Fail auth user LOGIN")
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: err.Error()})
		return
	}

	h.logger.Info("SUCCESS LOGIN", "user", req.Email)
	c.JSON(http.StatusOK, dto.Token(token))

}

func (h *AuthHandler) HandleDummy(c *gin.Context) {
	var req dto.DummyRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Error("DummyLogin bad req")
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: err.Error()})
		return
	}

	// Валидация role
	if !validations.IsValidRole(req.Role) {
		h.logger.Error("Fail role validations")
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: "Invalid role"})
		return
	}

	token, err := h.authService.DummyToken(c.Request.Context(), req.Role)
	if err != nil {
		h.logger.Error("Fail dummy token")
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: err.Error()})
		return
	}

	h.logger.Info("SUCCESS Dummy")
	c.JSON(http.StatusOK, dto.Token(token))
}

func (h *AuthHandler) HandleRegister(c *gin.Context) {
	var req dto.RegisterRequest
	var resp dto.RegisterResponse

	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Error("Register bad req")
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: err.Error()})
		return
	}

	// Валидация email
	if !validations.IsValidEmail(req.Email) {
		h.logger.Error("Fail email validations")
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: "Invalid email"})
		return
	}

	// Валидация role
	if !validations.IsValidRole(req.Role) {
		h.logger.Error("Fail role validations")
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: "Invalid role"})
		return
	}

	// Валидация pass
	if !validations.IsValidPassword(req.Password) {
		h.logger.Error("Fail pass validations")
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: "Invalid pass"})
		return
	}

	hpass, err := middleware.GeneratePasswordHash(req.Password)
	if err != nil {
		h.logger.Error("failed to HASH password REGISTER")
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: "Fail to hash password"})
		return
	}

	// маппим в модель
	model := mapper.RegisterRequestToUser(req, hpass)

	m, err := h.authService.CreateUser(c.Request.Context(), model)
	if err != nil {
		h.logger.Error("Failed to create user", "user", req.Email)
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: "User already exist"})
		return
	}

	resp = mapper.UserToRegisterResponse(m)

	h.logger.Info("SUCCESS REGISTER", "user", req.Email)
	c.JSON(http.StatusCreated, resp)
}
