package mapper

import (
	"log/slog"
	"pvz/internal/handler/dto"
	"pvz/internal/model"

	"github.com/google/uuid"
)

func ReceptionReqToModel(req dto.ReceptionRequst) model.Reception {
	id, err := uuid.Parse(req.PvzID)
	if err != nil {
		slog.Warn("Failed to parse pvzID", "error", err)
		return model.Reception{}
	}
	return model.Reception{
		PvzId: id,
	}
}

func ModelToReceptionResponse(model model.Reception) dto.ReceptionResponse {
	return dto.ReceptionResponse{
		Id:       model.Id.String(),
		DateTime: model.DateTime.String(),
		PvzID:    model.PvzId.String(),
		Status:   model.Status,
	}
}

func ReceptionReqToPvz(req dto.ReceptionRequst) model.Pvz {
	id, err := uuid.Parse(req.PvzID)
	if err != nil {
		slog.Warn("Failed to parse pvzID", "error", err)
		return model.Pvz{}
	}
	return model.Pvz{
		Id: id,
	}
}
