package entities

import (
	"pvz/internal/model"
	"time"

	"github.com/google/uuid"
)

type Reception struct {
	Id       uuid.UUID `db:"id"`
	DateTime time.Time `db:"date_time"`
	PvzId    uuid.UUID `db:"pvz_id"`
	Status   string    `db:"status"`
}

func ReceptionToEntity(model model.Reception) Reception {
	return Reception{
		Id:       model.Id,
		DateTime: model.DateTime,
		PvzId:    model.PvzId,
		Status:   model.Status,
	}
}

func EntityToReception(r Reception) model.Reception {
	return model.Reception{
		Id:       r.Id,
		DateTime: r.DateTime,
		PvzId:    r.PvzId,
		Status:   r.Status,
	}
}
