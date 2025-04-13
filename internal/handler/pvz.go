package handler

import (
	"context"
	"log/slog"
	"net/http"
	"pvz/internal/handler/dto"
	"pvz/internal/mapper"
	"pvz/internal/middleware/validations"
	"pvz/internal/model"

	"github.com/gin-gonic/gin"
)

type PvzService interface {
	CreatePvz(ctx context.Context, pvz model.Pvz) (model.Pvz, error)
}

type PvzHandler struct {
	pvzService PvzService
	logger     *slog.Logger
}

func NewPvzHandler(as PvzService, logger *slog.Logger) *PvzHandler {
	return &PvzHandler{pvzService: as, logger: logger}
}

func (h *PvzHandler) HandleCreatePvz(c *gin.Context) {
	var req dto.PvzCreateRequest
	var resp dto.PvzCreateResponse
	var err error

	if err = c.ShouldBindJSON(&req); err != nil {
		h.logger.Error("Invalid requst CreatePVZ")
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: "Invalid requst CreatePVZ"})
		return
	}

	// валидация города
	if !validations.IsValidCity(req.City) {
		h.logger.Error("Failed city validations", "city", req.City)
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: "Incorrect city! Legal: Казань, Москва, Санкт-Петербург"})
		return
	}

	model := mapper.PvzCreateRequestToPvz(req)
	model, err = h.pvzService.CreatePvz(c.Request.Context(), model)
	if err != nil {
		h.logger.Error("Failed create pvz")
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: err.Error()})
	}

	resp = mapper.PvzToPvzCreateResponse(model)

	h.logger.Info("SUCCESS CreatePVZ", "id", resp.Id)
	c.JSON(http.StatusCreated, resp)

}
