package entities

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID `db:"id"`
	Email    string    `db:"email"`
	Role     string    `db:"role"`
	Password string    `db:"password"`
}
