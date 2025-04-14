package mapper

import (
	"log/slog"
	"pvz/internal/handler/dto"
	"pvz/internal/model"

	"github.com/google/uuid"
)

func PvzCreateRequestToPvz(req dto.PvzCreateRequest) model.Pvz {
	return model.Pvz{
		City: req.City,
	}
}

func PvzToPvzCreateResponse(pvz model.Pvz) dto.PvzCreateResponse {
	return dto.PvzCreateResponse{
		Id:      pvz.Id.String(),
		RegDate: pvz.RegistrationDate.String(),
		City:    pvz.City,
	}
}

func CreateProductReqToPvz(req dto.CreateProductReq) model.Pvz {
	id, err := uuid.Parse(req.PvzId)
	if err != nil {
		slog.Warn("Failed to parse pvzID", "error", err)
		return model.Pvz{}
	}
	return model.Pvz{
		Id: id,
	}
}
