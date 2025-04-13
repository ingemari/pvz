package entities

import (
	"pvz/internal/model"
	"time"

	"github.com/google/uuid"
)

type Pvz struct {
	Id               uuid.UUID `db:"id"`
	RegistrationDate time.Time `db:"registration_date"`
	City             string    `db:"city"`
}

func PvzToEntity(pvz model.Pvz) Pvz {
	return Pvz{
		Id:               pvz.Id,
		RegistrationDate: pvz.RegistrationDate,
		City:             pvz.City,
	}
}

func EntityToPvz(ent Pvz) model.Pvz {
	return model.Pvz{
		Id:               ent.Id,
		RegistrationDate: ent.RegistrationDate,
		City:             ent.City,
	}
}
