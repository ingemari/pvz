package mapper

import (
	"pvz/internal/handler/dto"
	"pvz/internal/model"
)

func CreateProductReqToProduct(req dto.CreateProductReq) model.Product {
	return model.Product{
		Type: req.Type,
	}
}

func ProductToCreateProductResp(product model.Product) dto.CreateProductResp {
	return dto.CreateProductResp{
		Id:          product.Id.String(),
		DateTime:    product.DateTime.String(),
		Type:        product.Type,
		ReceptionId: product.ReceptionId.String(),
	}
}
