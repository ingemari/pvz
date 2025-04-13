package entities

import (
	"log/slog"
	"pvz/internal/model"

	"github.com/google/uuid"
)

type User struct {
	Id       uuid.UUID `db:"id"`
	Email    string    `db:"email"`
	Role     string    `db:"role"`
	Password string    `db:"password"`
}

func UserToEntity(user model.User) User {
	var id uuid.UUID
	var err error

	if user.Id == "" {
		id = uuid.New()
	} else {
		id, err = uuid.Parse(user.Id)
		if err != nil {
			slog.Error("parse UUID", "err", err)
			id = uuid.New()
		}
	}

	return User{
		Id:       id,
		Email:    user.Email,
		Role:     user.Role,
		Password: user.Password,
	}
}

func EntityToUser(ent User) model.User {
	return model.User{
		Id:       ent.Id.String(),
		Email:    ent.Email,
		Role:     ent.Role,
		Password: ent.Password,
	}
}
