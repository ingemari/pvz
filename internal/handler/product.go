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

type ProductService interface {
	CreateProduct(ctx context.Context, product model.Product, pvz model.Pvz) (model.Product, error)
}

type ProductHandler struct {
	productService ProductService
	logger         *slog.Logger
}

func NewProductHandler(as ProductService, logger *slog.Logger) *ProductHandler {
	return &ProductHandler{productService: as, logger: logger}
}

func (h *ProductHandler) HandleCreateProduct(c *gin.Context) {
	var req dto.CreateProductReq
	var err error

	if err = c.ShouldBindJSON(&req); err != nil {
		h.logger.Error("Invalid requst CreatePVZ")
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: "Invalid requst CreatePVZ"})
		return
	}

	pvz := mapper.CreateProductReqToPvz(req)
	product := mapper.CreateProductReqToProduct(req)

	result, err := h.productService.CreateProduct(c.Request.Context(), product, pvz)
	if err != nil {
		h.logger.Error("Failed to create product", "err", err)
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: "Failed to create product"})
		return
	}

	resp := mapper.ProductToCreateProductResp(result)

	c.JSON(http.StatusCreated, resp)
}
