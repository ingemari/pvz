package mapper

import (
	"pvz/internal/handler/dto"
	"pvz/internal/model"
	"pvz/internal/repository/entities"
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

func PvzToEntity(pvz model.Pvz) entities.Pvz {
	return entities.Pvz{
		Id:               pvz.Id,
		RegistrationDate: pvz.RegistrationDate,
		City:             pvz.City,
	}
}

func EntityToPvz(ent entities.Pvz) model.Pvz {
	return model.Pvz{
		Id:               ent.Id,
		RegistrationDate: ent.RegistrationDate,
		City:             ent.City,
	}
}
