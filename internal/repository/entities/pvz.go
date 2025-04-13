package entities

import (
	"time"

	"github.com/google/uuid"
)

type Pvz struct {
	Id               uuid.UUID `db:"id"`
	RegistrationDate time.Time `db:"registration_date"`
	City             string    `db:"city"`
}
