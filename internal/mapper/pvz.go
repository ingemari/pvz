package mapper

import (
	"pvz/internal/handler/dto"
	"pvz/internal/model"
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
