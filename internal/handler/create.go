package handler

import (
	"context"
	"net/http"
	"pvz/internal/model"

	"github.com/gin-gonic/gin"
)

// type CreateUserRequest struct {
// 	Username string `json:"username"`
// }

// type CreateUserResponse struct {
// 	Id       string `json:"id"`
// 	Username string `json:"username"`
// }

type AuthService interface {
	Register(ctx context.Context, req model.RegisterRequest) (*model.RegisterResponse, error)
}

type AuthHandler struct {
	authService AuthService
}

func NewAuthHandler(as AuthService) *AuthHandler {
	return &AuthHandler{authService: as}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req model.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	resp, err := h.authService.Register(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Registration failed"})
		return
	}

	c.JSON(http.StatusCreated, resp)
}
