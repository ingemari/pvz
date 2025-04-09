package models

import "time"

// Error defines model for Error.
type Error struct {
	Message string `json:"message"`
}

// PVZ defines model for PVZ.
type PVZ struct {
	City             PVZCity             `json:"city"`
	Id               *openapi_types.UUID `json:"id,omitempty"`
	RegistrationDate *time.Time          `json:"registrationDate,omitempty"`
}

// Product defines model for Product.
type Product struct {
	DateTime    *time.Time          `json:"dateTime,omitempty"`
	Id          *openapi_types.UUID `json:"id,omitempty"`
	ReceptionId openapi_types.UUID  `json:"receptionId"`
	Type        ProductType         `json:"type"`
}

// Reception defines model for Reception.
type Reception struct {
	DateTime time.Time       `json:"dateTime"`
	Id       uuid.UUID       `json:"id,omitempty"`
	PvzId    uuid.UUID       `json:"pvzId"`
	Status   ReceptionStatus `json:"status"`
}

// Token defines model for Token.
type Token = string

// User defines model for User.
type User struct {
	Email openapi_types.Email `json:"email"`
	Id    *openapi_types.UUID `json:"id,omitempty"`
	Role  UserRole            `json:"role"`
}
