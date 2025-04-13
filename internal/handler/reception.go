package handler

import (
	"context"
	"log/slog"
	"net/http"
	"pvz/internal/handler/dto"
	"pvz/internal/mapper"
	"pvz/internal/model"

	"github.com/gin-gonic/gin"
)

type ReceptionService interface {
	CreateReception(ctx context.Context, reception model.Reception, pvz model.Pvz) (model.Reception, error)
}

type ReceptionHandler struct {
	receptionService ReceptionService
	logger           *slog.Logger
}

func NewReceptionHandler(h ReceptionService, logger *slog.Logger) *ReceptionHandler {
	return &ReceptionHandler{receptionService: h, logger: logger}
}

func (h *ReceptionHandler) HandleReceptionCreate(c *gin.Context) {
	var req dto.ReceptionRequst
	var resp dto.ReceptionResponse

	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Error("Invalid requst CreatePVZ")
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: "Invalid request Reception create"})
		return
	}

	reception := mapper.ReceptionReqToModel(req)

	pvz := mapper.ReceptionReqToPvz(req)
	m, err := h.receptionService.CreateReception(c.Request.Context(), reception, pvz)
	if err != nil {
		h.logger.Error("Failed create reception")
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: err.Error()})
	}

	resp = mapper.ModelToReceptionResponse(m)
	c.JSON(http.StatusCreated, resp)
}
