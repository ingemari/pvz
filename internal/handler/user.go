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
}

func NewAuthHandler(as AuthService) *AuthHandler {
	return &AuthHandler{authService: as}
}

func (h *AuthHandler) HandleLogin(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: err.Error()})
		return
	}

	// Валидация email
	if !validations.IsValidEmail(req.Email) {
		slog.Error("Fail email validations")
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: "Invalid email format"})
		return
	}

	// mapping in model
	model := mapper.LoginRequestToUser(req, []byte(req.Password))

	token, err := h.authService.LoginUser(c.Request.Context(), model)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.Token(token))

}

func (h *AuthHandler) HandleDummy(c *gin.Context) {
	var req dto.DummyRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: err.Error()})
		return
	}

	// Валидация role
	if !validations.IsValidRole(req.Role) {
		slog.Error("Fail role validations")
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: "Invalid role"})
		return
	}

	token, err := h.authService.DummyToken(c.Request.Context(), req.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.Token(token))

}

func (h *AuthHandler) HandleRegister(c *gin.Context) {
	var req dto.RegisterRequest
	var resp dto.RegisterResponse

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: err.Error()})
		return
	}

	hpass, err := middleware.GeneratePasswordHash(req.Password)
	if err != nil {
		slog.Warn("failed to hash password", "error", err)
		return
	}

	// Валидация email
	if !validations.IsValidEmail(req.Email) {
		slog.Error("Fail email validations")
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: "Invalid email format"})
		return
	}

	// Валидация role
	if !validations.IsValidRole(req.Role) {
		slog.Error("Fail role validations")
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: "Invalid role"})
		return
	}

	// маппим в модель
	model := mapper.RegisterRequestToUser(req, hpass)

	m, err := h.authService.CreateUser(c.Request.Context(), model)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: err.Error()})
		return
	}

	resp = mapper.UserToRegisterResponse(m)

	c.JSON(http.StatusOK, resp)
}
